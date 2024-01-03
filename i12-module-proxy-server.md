### Module Proxy Servers

- **Every Go module is stored in a source code repository, like GitHub or GitLab. But by default, go get doesn’t fetch code directly from source code repositories**. 

- **Instead, it sends requests to a proxy server run by Google. This server keeps copies of every version of virtually all public Go modules**. 

- **If a module or a version of a module isn’t present on the proxy server, it downloads the module from the module’s repository, stores a copy, and returns the module**.

- **In addition to the proxy server, Google also maintains a sum database. It stores information on every version of every module. This includes the entries that appear in a go.sum file for the module at that version and a signed, encoded tree description that contains the record**. 

- Just as **the proxy server protects you from a module or a version of a module being removed from the internet, the sum database protects you against modifications to a version of a module. This could be malicious**.

- **Every time you download a module via go build, go test, or go get, the Go tools calculate a hash for the module and contact the sum database to compare the calculated hash to the hash stored for that module’s version. If they don’t match, the module isn’t installed**.


### Specifying a Proxy Server

- Some people object to sending requests for third-party libraries to Google. There are a few options:
  
- If you don’t mind a public proxy server, but don’t want to use Google’s, you can
switch to **GoCenter (which is run by JFrog) by setting the GOPROXY environment variable to https://gocenter.io,direct**.

- You can **disable proxying** entirely by **setting the GOPROXY environment variable to direct** You’ll download modules directly from their repositories, but if you depend on a version that’s removed from the repository, you won’t be able to access it.

- You **can run your own proxy server. Both Artifactory and Sonatype have Go proxy server support built into their enterprise repository products**. 

- **The Athens Project provides an open source proxy server. Install one of these products on your network and then point GOPROXY to the URL**.

- If you are using a **public proxy server, you can set the GOPRIVATE environment variable to a comma-separated list of your private repositories**. For example, **if you set GOPRIVATE to**:

```sh
    GOPRIVATE=*.example.com,company.com/repo
```

- **Any module stored in a repository that’s located at any subdomain of example.com or at a URL that starts with company.com/repo will be downloaded directly**.