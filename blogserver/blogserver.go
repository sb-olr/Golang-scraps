// blogserver (experiment)
package blogserver

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Comment struct {
	Title  string
	Body   []byte
	Date   string
	Author string
	Email  string
}

type Post struct {
	Title    string
	Body     []byte
	Date     string
	Comments []Comment // get following comments
}

// create post
func (p *Post) create() error {
	// create file for storing the post info
	_, err := os.Create("content/posts/" + strings.Replace(p.Title, " ", "-", -1) + ".txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// enter the post into the file
	return ioutil.WriteFile("content/posts/"+strings.Replace(p.Title, " ", "-", -1)+".txt", p.Body, 0600)
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
	lines, err := readLines("/home/ndegree/Desktop/GoApp/blogserver/blog.conf")
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
func getPost(title string) (post Post, error error) {
	post.Title = strings.Replace(title, "-", " ", -1)
	filename := title + ".txt"
	post.Body, error = ioutil.ReadFile("content/posts/" + filename)  // get body of the post
	commentList, error := ReadDir("content/comments/" + title + "/") // get the comment following
	//var tempc []Comment
	for _, comment := range commentList {
		var c Comment
		c.Title = strings.Replace(strings.Replace(comment.Name(), "-", " ", -1), ".txt", " ", -1)
		c.Body, error = ioutil.ReadFile("content/comments/" + title + "/" + comment.Name())
		post.Comments = append(post.Comments, c)
	}
	return post, nil
}

const (
	viewLen = len("/view/")
)

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[viewLen:] // get post name from url
	post, err := getPost(title)
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	t, err := template.ParseFiles("content/view.html") // get the template content from the html file
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post.Body)
	t.Execute(w, post) // execute the template to show the website

	// if input comment
	c := r.FormValue("commenting")
	_, err = os.Create("content/comments/" + strings.Replace(title, " ", "-", -1) + "/" + strings.Replace(c, " ", "-", -1) + ".txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// enter the post into the file
	ioutil.WriteFile("content/comments/"+strings.Replace(title, " ", "-", -1)+"/"+strings.Replace(c, " ", "-", -1)+".txt", []byte(c), 0600)
}

func new(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "content/new.html") // the interface to get the post input
}

func create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	p := &Post{Title: title, Body: []byte(body)}
	p.create()
	http.Redirect(w, r, "/view/"+strings.Replace(title, " ", "-", -1), http.StatusFound) // view the post
}

// handle the root
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func comment(w http.ResponseWriter, r *http.Request) {
	// if input comment
	c := r.FormValue("commenting")
	title := r.URL.Path[viewLen:] // ge post name from url
	_, err := os.Create("content/comments/" + strings.Replace(title, " ", "-", -1) + strings.Replace(c, " ", "-", -1) + ".txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// enter the post into the file
	ioutil.WriteFile("content/comments/"+strings.Replace(title, " ", "-", -1)+"/"+strings.Replace(c, " ", "-", -1)+".txt", []byte(c), 0600)
	http.Redirect(w, r, "/view/"+strings.Replace(title, " ", "-", -1)+"/", http.StatusFound) // view the post
}

func init() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/view/", view)
	http.HandleFunc("/new/", new)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/comment/", comment)
}
