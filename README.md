# goScRp
A Command Line Interfaced Quotes webscrapper written in go.

This uses the go-colly package to scrape quotes from [quotes.toscrape.com](https://quotes.toscrape.com/) website.

## Installation

To Modify the code :
```
git clone git@github.com:2k4sm/goScRp.git
cd goScRp
```

## Usage

To generate a random quote:

```
./goScRp
```

To generate quotes based on `tags`:

```
./goScRp -t `<tag>`
```

To generate quotes based on a given `page`:

```
./goScRp -p `<page>`
```

To generate quotes based on `tags` and which is present on a given `page`:

```
./goScRp -t `<tag>` -p `<page>`
```

To generate details about an `author`:

```
./goScRp -a `<author-name>`
```



# License

    MIT License

    Copyright (c) 2023 Shrinibas Mahanta
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.