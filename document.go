package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Document struct {
	*Client
}

type DocumentOption struct {
	Slug  string
	Title string
	Body  string
	/*
		"markdown", "lake", "html"
		default is markdown
	*/
	Format string
	/*
		0 - private
		1 - public
	*/
	Public int
	/*
		If you edit the document on the page, then the document will turn into a Lake format,
		if you can use MarkDown to update, this is what you need to add _force_asl = 1 to
		ensure the correct conversion of the content.
	*/
	ForceASL int
}

func newDocument(c *Client) *Document {
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
func (d *Document) GetInfo(namespace, slug string) (*internal.ResponseDocDetailSerializer, error) {
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

// Create create a document
func (d *Document) Create(namespace string, options ...DocumentOption) (*internal.ResponseDocDetailSerializer, error) {
	var opt DocumentOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentCreatePath, namespace)
		document = internal.ResponseDocDetailSerializer{}
		body     = struct {
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			Body   string `json:"body"`
			Format string `json:"format"`
			Public int    `json:"public"`
		}{
			Slug:   opt.Slug,
			Title:  opt.Title,
			Body:   opt.Body,
			Format: opt.Format,
			Public: opt.Public,
		}
	)

	respBody, err := d.do(url, option{Method: "POST", Body: body})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (d *Document) getDocumentID(namespace, slug string) (int, error) {
	doc, err := d.GetInfo(namespace, slug)
	return doc.Data.ID, err
}

// Update update a document
func (d *Document) Update(namespace, slug string, options ...DocumentOption) (*internal.ResponseDocDetailSerializer, error) {
	var opt DocumentOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	documentID, err := d.getDocumentID(namespace, slug)
	if err != nil {
		return nil, ErrGetDocumentID
	}
	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentUpdatePath, namespace, documentID)
		document = internal.ResponseDocDetailSerializer{}
		body     = struct {
			Slug     string `json:"slug"`
			Title    string `json:"title"`
			Body     string `json:"body"`
			Public   int    `json:"public"`
			ForceASL int    `json:"_force_asl"`
		}{
			Slug:     opt.Slug,
			Title:    opt.Title,
			Body:     opt.Body,
			Public:   opt.Public,
			ForceASL: opt.ForceASL,
		}
	)

	respBody, err := d.do(url, option{Method: "PUT", Body: body})
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
func (d *Document) Delete(namespace, slug string) (*internal.ResponseDocDetailSerializer, error) {
	documentID, err := d.getDocumentID(namespace, slug)
	if err != nil {
		return nil, ErrGetDocumentID
	}

	var (
		url      = fmt.Sprintf(d.BaseURL+internal.DocumentDeletePath, namespace, documentID)
		document = internal.ResponseDocDetailSerializer{}
	)

	respBody, err := d.do(url, option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
