// blogserver (experiment)
package blogserver

import (
	"appengine"           // for logging in the appengine
	"appengine/datastore" // a database
	"appengine/user"      // and return to the handler
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Comment struct {
	Title  string
	Body   []byte
	Time   time.Time
	Author string
	Email  string
}

type Post struct {
	Author   string
	Title    string
	Body     []byte
	Time     time.Time
	Comments []Comment // get following comments
}

type byDate []os.FileInfo

func (f byDate) Len() int           { return len(f) }
func (f byDate) Less(i, j int) bool { return time.Since(f[i].ModTime()) > time.Since(f[j].ModTime()) }
func (f byDate) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// get files form given dirname
func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Sort(byDate(list))
	return list, nil
}

// read lines from a file
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		// if it has reached new line
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func getConfigValue(key string) (value string) {
	lines, err := readLines("blogserver/blog.conf")
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	for _, line := range lines {
		configData := strings.Split(line, ":")
		if len(configData) > 2 {
			fmt.Println("Check ConfigFile, and option has more than value")
			return
		}
		if configData[0] == key {
			value = configData[1]
			break
		}
	}
	return value
}

// get the root of the website with the config file
var wwwroot string = getConfigValue(("wwwroot"))

// get content and info of the posts and the comments following
func getPost(w http.ResponseWriter, r *http.Request, title string) (post Post, error error) {
	var p []Post
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Post").Filter("Title =", title).Order("-Time").Limit(1)
	if _, err := q.GetAll(c, &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post = Post(p[0])
	return post, nil
}

const (
	viewLen = len("/view/")
)

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[viewLen:] // get post name from url
	post, err := getPost(w, r, title)

	t, err := template.ParseFiles("blogserver/content/view.html") // get the template content from the html file
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post.Body)
	t.Execute(w, post) // execute the template to show the website

	commenting := r.FormValue("commenting")
	if commenting != "" {
		comment(w, r, commenting, &post)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Post").Order("-Time")
	if _, err := q.GetAll(c, &posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("blogserver/content/list.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, posts) // execute the template to show the website
}

func new(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "blogserver/content/new.html") // the interface to get the post input
}

func create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	p := Post{Time: time.Now(), Title: title, Body: []byte(body)}
	c := appengine.NewContext(r)
	if u := user.Current(c); u != nil {
		p.Author = u.String()
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Post", nil), &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // view the post
}

// handle the root
func index_handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	t, err := template.ParseFiles("blogserver/content/Index.html") // get the template content from the html file
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
	t.Execute(w, u) // execute the template to show the website
}

func comment(w http.ResponseWriter, r *http.Request, commenting string, p *Post) {
	comment := Comment{Time: time.Now(), Body: []byte(commenting)}

	c := appengine.NewContext(r)
	if u := user.Current(c); u != nil {
		comment.Author = u.String()
	}
	p.Comments = append(p.Comments, comment)
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Post", nil), &p)

	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/view/"+p.Title, http.StatusFound) // view the comment
}

func init() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/view/", view)
	http.HandleFunc("/list", list)
	http.HandleFunc("/new/", new)
	http.HandleFunc("/create/", create)
}
