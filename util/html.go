package util

import (
    "html/template"
    "io"
)

func ExecuteTemplate(w io.Writer, data any, filenames... string) error {
    var t *template.Template
    var err error

    if t, err = template.ParseFiles(filenames...); err != nil {
        return err
    }

    if err = t.Execute(w, data); err != nil {
        return err
    }

    return nil
}
