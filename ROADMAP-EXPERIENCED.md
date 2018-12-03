# ROADMAP FOR EXPERIENCED DEVELOPERS

Hi!

If you're an experienced developer, you might want to follow this roadmap while taking this course.

This course starts from the most basics than advances toward the end, step by step. So, the complexity of the topics increases on each step. I've intentionally designed it so to make it easy for everyone. 

If you think some of the topics are easy for you, then watch the recap lectures and skip the lectures in that section altogether, you can always come back to them later.

BTW, if you're also wondering about why should you learn Go, then scroll down to the bottom of this article.

---

## Basic Configuration

1. **Git clone the repo**
   https://github.com/inancgumus/learngo

2. Read **"Side Note About Go Modules"** part right after the outline below.
   1. If you also want to learn about GOPATH, watch: "Learn about GOPATH and Go directory structure" lecture

3. You might want to **increase the video speed**.
    1. If you don't know how, [here](https://support.udemy.com/hc/en-us/articles/229231247-Change-the-Video-Speed)'s how you can do that.

---

## Roadmap for Experienced Developers

👉**Then, watch the following lectures in this order.**

### Getting Started!
* Install Go
* Configure: Visual Studio Code

### PART I — Write Your First Go Program
* Code your first program
* Compile and Run your first program using Go Build
* Run your first program using Go Run

### PART I — Packages, Scopes and Importing
* Packages - Learn how to run multiple files
* Packages - Executable vs. Library Packages
* Importing and File Scope
* Importing - Rename imported packages

### PART I — Statements, Expressions, and Comments
* Go Doc: Generate documentation automatically from your code

### PART I — Create Your First Library Package
* Watch all the lectures here.

### PART II — Variables and Type Conversion
* Zero-Values
* Variable Declaration Examples (Code Along)
* Example: Greeter: Get input from the command-line
* Short Declaration: Initialization & Type Inference
* Short Declaration: Package Scope
* Redeclaration
* When to use short declaration?
* Recap
* Naming Things: Recommendations

### PART II — Printing Formatted Output using Printf
* Printf: Recap: Let's summarize

### PART II — Numbers and Strings
* Numbers: IncDec: Easy Way to Increase and Decrease
* Strings: Raw String Literals
* Strings: Get the length of a string and Introduction to Runes

### PART II — Deeper Into The Go Type System
* Watch all the lectures here.

### Constants
* Recap: Constants
* Typeless Constants: Understanding Typelessness
* Typeless Constants: Default Types
* Example: Real-Life Usage
* IOTA: Constant Number Generator

### PART III — If Statement and Error Handling
* Recap: If Statement
* Then watch all the lectures starting with:
    * Error Handling Basics: Introduction

### PART III — Switch Statement
* Fallthrough Statement
* Recap

### PART III — Loops
* Continue Statement and Debugging with Delve
* For Statement: Looping over Slices
* For Range Clause
* Recap: Loops
* Project: Lucky Number Part I: Randomization
* Project: Lucky Number Part II: Seeding with time
* Project: Lucky Number Part III
* Labeled Statements: Labeled Break and Continue
* Labeled Statements: Break from a Switch
* Labeled Statements: Goto (Optional)

The rest is coming soon.

---

## SIDE-NOTE ABOUT GO MODULES

Since Go 1.11, there's a new feature called Go Modules (_it was known as vgo before_).

With Go Modules, you are not required to put your source code files under $GOPATH anymore. You can put them in any directory that you want.

If you're a power user and want to try this **experimental feature**, just put your programs in a directory outside of `$GOPATH` (this is very important). And then run this to initialize your module:

    `go mod init [module-name]`

Afterward, while in this directory, Go assumes that it's like you're under your GOPATH. However, you'll be under your module's directory. You can create as many modules as you want. However, keep it simple and use only one module for now.

### EXAMPLE:

1. Let's say that my module directory is under: `~/Desktop/go`
2. Also, my github.com username is: `inancgumus`
3. And, I have a program called: `hello`
4. So, I put my hello program under: `~/Desktop/go/hello`
5. Then, while under `~/Desktop/go` I would call mod init like this:

    `go mod init github.com/inancgumus`
    
6. So, if I want to build my program, Go into `~/Desktop/go/hello`
7. Moreover, I can use usual go build and go run commands there.
8. By doing so, I've created a single module for all my projects. As I've said, you can also create separate modules for each of your projects, like this:

    `go mod init github.com/inancgumus/hello`

For more information check out [this wiki article](https://github.com/golang/go/wiki/Modules).

---

## Why you should learn Go?

If you're curious about why you should learn Go, then check out this summary here.

**In summary:** Go is easy as Python and Javascript and it's as fast as C/C++. It's more enjoyable to work with Go than C/C++. You can go low-level, or you can stay high-level.

### WHAT GO IS USED FOR?

Go is used mostly by web companies: Google, Facebook, Twitter, Uber, Apple, Dropbox, Soundcloud, Medium, Mozilla Firefox, Github, Docker, Kubernetes, and Heroku.

**Go is best for:** Cross-Platform Command-line Tools, Distributed Network Applications, Cloud technologies like Microservices and Serverless, Web APIs, Database Engines, Big-Data Processing Pipelines, Embedded Development, and so on.

**Go is not best for (but doable):** Desktop Apps, Writing Operating Systems, Kernel Drivers, Game Development, etc

### WHO DESIGNED GO?

Go designed by one of the most influential people in the industry:

* Unix: Ken Thompson
* UTF-8, Plan 9: Rob Pike
* Hotspot JVM (Java Virtual Machine): Robert Griesemer

### WHAT YOU CAN DO WITH GO?

* [Network Driver written in Go](https://www.net.in.tum.de/fileadmin/bibtex/publications/theses/2018-ixy-go.pdf) (_only 10% penalty compared to C driver_)
* [Google gVisor](https://cloud.google.com/blog/products/gcp/open-sourcing-gvisor-a-sandboxed-container-runtime) (_User space kernel written in Go_)
* [Multi-platform Nintendo emulator](https://humpheh.github.io/goboy/)
* [Docker: Container system](https://github.com/moby/moby)
* [Kubernetes: Container scheduling and management](https://github.com/kubernetes/kubernetes)
* VM image deduplication utility
* Chat server
* RUM beacon collector
* Time-series database engine, a client for it, cmd line tools, etc
* Map-reduce library
* Clustered front-end-optimizing reverse proxy with on the fly content rewriting, image resizing, caching, Lua event handler execution (all multi-tenant)
* Geographically distributed reverse proxy CDN nodes
* Health management daemon with event handlers and peer to peer reporting
* Pure Go DNS server
* API backend that interfaces with MySQL
* Linux process capture/restore utility
* Reverse Proxy to mask our asset server from clients.
* HTML -> PDF converter for invoice generation.
* URL shortener like tinyurl.com and goo.gl
* SMS messaging service.
* Credit Card payment gateway
* JSON Web Token package
* On the fly image processing services
* 3d render farm/content production pipeline (pretty large project)
* Production lxc container deployment
* Automated testing management

Reference: [This reddit post](https://www.reddit.com/r/golang/comments/5nac2b/what_have_you_used_go_for_in_your_professional/).

### [From Eight years of Go post](https://blog.golang.org/8years):

> Today, **every single cloud company has critical components of their cloud infrastructure implemented in Go** including Google Cloud, AWS, Microsoft Azure, Heroku, and many others. Go is a vital part of cloud companies like Alibaba, Cloudflare, and Dropbox. Go is a critical part of open infrastructure including Kubernetes, Cloud Foundry, Openshift, NATS, Docker, Istio, Etcd, Consul, Juju and many more. Companies are increasingly choosing Go, to build cloud infrastructure solutions.

### HOW MUCH YOU CAN EARN?

+ [Go Salaries in US](https://www.payscale.com/research/US/Skill=Go_(Golang)_Programming_Language/Salary)

### Check out these posts for more information:

* [About Go: An Overview](https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c)
* [Why you should learn Go?](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)
* [Emerging language of cloud Infrastructure](https://redmonk.com/dberkholz/2014/03/18/go-the-emerging-language-of-cloud-infrastructure/)
* [Companies using Go](https://github.com/golang/go/wiki/GoUsers)
* [Eight years of Go](https://blog.golang.org/8years)
* [Twitter: Handling Five Billion Session in a Day with Go](https://blog.twitter.com/engineering/en_us/a/2015/handling-five-billion-sessions-a-day-in-real-time.html)
* [A C++ developer looks at Go](https://www.murrayc.com/permalink/2017/06/26/a-c-developer-looks-at-go-the-programming-language-part-1-simple-features/)
