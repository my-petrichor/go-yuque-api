package yuque

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Document struct {
	*client
}

func newDocument(c *client) *Document {
	return &Document{
		c,
	}
}

// ListAll list all documents in a repo
func (d *Document) ListAll(namespace string) (*internal.ResponseDocSerializer, error) {
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentListPath, namespace)
		document = internal.ResponseDocSerializer{}
	)

	respBody, err := d.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// Get get document detail info
func (d *Document) Get(namespace, slug string) (*internal.ResponseDocDetailSerializer, error) {
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentGetPath, namespace, slug)
		document = internal.ResponseDocDetailSerializer{}
	)

	respBody, err := d.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (d *Document) GetDocumentID(namespace, slug string) (int, error) {
	doc, err := d.Get(namespace, slug)
	return doc.Data.ID, err
}

// Create create a document
/*
 public
 0 - private
 1 - public
*/
/*
 format
 markdown, lake, html
 default is markdown
*/
func (d *Document) Create(namespace, newSlug, newTitle, format, body string, public int) (*internal.ResponseDocDetailSerializer, error) {
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentCreatePath, namespace, public)
		document = internal.ResponseDocDetailSerializer{}
	)

	respBody, err := d.do(url, Option{Method: "POST", Body: map[string]string{
		"title":  newTitle,
		"slug":   newSlug,
		"format": format,
		"body":   body,
	}})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// Update update a document
/*
 public
 0 - private
 1 - public
*/
/*
 If you edit the document on the page, then the document will turn into a Lake format,
 if you can use MarkDown to update, this is what you need to add _force_asl = 1 to
 ensure the correct conversion of the content.
*/
func (d *Document) Update(namespace, newTitle, newSlug, body string, id, public int, forceASIs ...int) (*internal.ResponseDocDetailSerializer, error) {
	var forceASI int
	if len(forceASIs) > 1 {
		return nil, errors.New("forceASI length more than one")
	} else if len(forceASIs) == 1 {
		forceASI = forceASIs[0]
	}
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentUpdatePath, namespace, id, public, forceASI)
		document = internal.ResponseDocDetailSerializer{}
	)

	respBody, err := d.do(url, Option{Method: "PUT", Body: map[string]string{
		"title": newTitle,
		"slug":  newSlug,
		"body":  body,
	}})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// Delete delete a document
func (d *Document) Delete(namespace string, documentID int) (*internal.ResponseDocDetailSerializer, error) {
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentDeletePath, namespace, documentID)
		document = internal.ResponseDocDetailSerializer{}
	)

	respBody, err := d.do(url, Option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
