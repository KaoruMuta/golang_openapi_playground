// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func decodeCreateTaskResponse(resp *http.Response) (res *CreateTaskCreated, _ error) {
	switch resp.StatusCode {
	case 201:
		// Code 201.
		return &CreateTaskCreated{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeDeleteTaskResponse(resp *http.Response) (res *DeleteTaskNoContent, _ error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &DeleteTaskNoContent{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetTasksResponse(resp *http.Response) (res []Task, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response []Task
			if err := func() error {
				response = make([]Task, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Task
					if err := elem.Decode(d); err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if response == nil {
					return errors.New("nil is invalid value")
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodePartiallyUpdateTaskResponse(resp *http.Response) (res *PartiallyUpdateTaskNoContent, _ error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PartiallyUpdateTaskNoContent{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUpdateTaskResponse(resp *http.Response) (res *UpdateTaskNoContent, _ error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &UpdateTaskNoContent{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
