# Shelf

A simple document management system.

```   
      _ _
 .-. | | |
 | |_|.|.|_
 | | | | | |<\
 |S|H|E|L|F| \\
 | | | | | |  \\
 | |$| | | |   \>
""""""""""""""""""
```
*Adapted "Bookshelf" by David S. Issel*

## Installation

If you have Go installed already

```
go get github.com/matthistuff/shelf
```

Shelf uses mongoDB for data and file storage. You need to expose your mongo host to Shelf.

```
export SHELF_DB_HOST=your_mongo:port
```

## Usage

Shelf is a CLI utility to manage documents. Documents are plain containers for attributes and files.

```
shelf create my awesome document
> 55b6325c2903461865000001
```

Add a attribute.

```
shelf attribute add 55b6325c2903461865000001 domain silly stuff
```

Create another document.

```
shelf create another awesome document
> 55b632d6290346187d000001
```

Attach a file.

```
shelf attach 55b632d6290346187d000001 ~/Documents/secret.pdf
> 55b63311290346188b000001
```

Get document info.

```
shelf info 55b632d6290346187d000001
>another awesome document
>
>Created at Mon, 27 Jul 2015 15:32:06 CEST
>
>Attributes
>
>Attachments
>	(1) 55b63311290346188b000001: secret.pdf (Mon, 27 Jul 2015 15:33:05 CEST)
```

Search for documents.

```
shelf search awesome
>(1) 55b6325c2903461865000001 "my awesome document"
>(2) 55b632d6290346187d000001 "another awesome document"
>Page 1 of 1
```

Special search.

```
shelf search domain:silly
>(1) 55b6325c2903461865000001 "my awesome document"
>Page 1 of 1
```

## Command reference

**create** *\<title>...
Create a new document

**delete** *\<document-id>*
Delete a document

**info** *\<document-id>*
Print information about an object

**list** *[--page=\<page>]*
Lists all available objects

**attach** *\<document-id> \<path-to-file>*
Attaches a file to a document

**attachments** *\<document-id>*
Lists all attachments of a document

**retrieve** *\<attachment-id>*
Sends an attachment to stdout

**attribute add** *\<document-id> \<attribute-name> \<attribute-value>...*
Adds an attribute to an object

**attribute remove** *\<document-id> \<attribute-name> \<attribute-value>...*
Removes an attribute from an object

**tag** *\<document-id> \<tag-value>*
Short for *attribute add \<document-id> tag \<tag-value>...*

**untag** *\<document-id> \<tag-value>*
Short for *attribute remove \<document-id> tag \<tag-value>...*

**Search** *\<query>...*
Search documents

Search does a full text search on the document title, attachment content (soon) and attribute values with different weighting.

You can limit the search to documents with certain attributes by using the syntax *\<attribute-name>:\<attribute-value>*.

Search always uses logical AND when combining text and attribute searches. The position of the items does not matter.

## Short Ids

All commands support using short ids from listed object outputs. Whenever you see a number in parenthesis next to an 24 character hex (`(1) 55b6325c2903461865000001 ...`), you can use the short number as an alias to the long hex id.