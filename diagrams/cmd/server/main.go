package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Flags struct {
	Addr string
	Root string
}

func (f *Flags) Register(fs *flag.FlagSet) {
	fs.StringVar(
		&f.Addr,
		"addr",
		":8099",
		"http address to bind")
	fs.StringVar(
		&f.Root,
		"root",
		".",
		"the directory to serve files from")
}

const index string = `
<!DOCTYPE html>
<html>
<head>
  	<title>pin</title>
  	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  	<link href='http://fonts.googleapis.com/css?family=Raleway:400,100,200' rel='stylesheet' type='text/css'>
	<style>
		#root {
			width: 900px;
			margin: 100px auto;
		}
		#root.lg .pos-sm {
			display: none;
		}
		#root.md {
			width: 600px;
		}
		#root.md .pos-lg {
			display: none;
		}
		#root.sm {
			width: 300px;
		}
		#root.sm .pos-lg {
			display: none;
		}
		#t {
			position: fixed;
			top: 20px;
			left: 20px;
			width: 15px;
			height: 15px;
			border-radius: 15px;
			background-color: #f6f6f6;
			border: 1px solid #eee;
			cursor: pointer;
		}
		#t:hover {
			background-color: #eee;
		}
		#t:active {
			background-color: #999;
		}
	</style>
</head>
<body>
	<div id="t"></div>
	<div id="root" class="lg">
		<p>
		Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
		</p>
		<p>
			<img src="pos-lg.svg" class="pos-lg">
			<img src="pos-sm.svg" class="pos-sm">
		</p>
		<p>
		Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
		</p>
		<p>
			<img src="pos-lg-summary.svg" class="pos-lg">
			<img src="pos-sm-summary.svg" class="pos-sm">
		</p>
		<p>
			Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
		</p>
		<p>
			<img src="neg-lg-summary.svg" class="pos-lg">
			<img src="neg-sm-summary.svg" class="pos-sm">
		</p>
		<p>
			Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
		</p>
	</div>
	<script>
		const root = document.querySelector('#root');
		document.querySelector('#t').addEventListener('click', () => {
			const cl = root.classList;
			if (cl.contains('lg')) {
				cl.remove('lg');
				cl.add('md');
			} else if (cl.contains('md')) {
				cl.remove('md');
				cl.add('sm');
			} else {
				cl.remove('sm');
				cl.add('lg');
			}
		});
	</script>
</body>
</html>
`

func newHandler(root string) http.Handler {
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				if _, err := fmt.Fprint(w, index); err != nil {
					log.Panic(err)
				}
			} else {
				fs.ServeHTTP(w, r)
			}
		})
}

func main() {
	var flags Flags
	flags.Register(flag.CommandLine)
	flag.Parse()

	log.Panic(http.ListenAndServe(
		flags.Addr,
		newHandler(flags.Root),
	))
}
