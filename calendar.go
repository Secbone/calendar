package main

import (
	"io"
	"os"
	"time"
	"bytes"
	"text/template"
)

type Holiday struct {
	Name string
	StartDate string
	EndDate string
	DTStamp string
	Rest bool
}

type Calendar struct {
	Name string
	Color string
	Holidays []Holiday
	t *template.Template
	prefix string
}

func NewCalendar(name string, color string) *Calendar {
	t := template.Must(template.New("iCalendar").ParseFiles("icalendar.tmpl"))

	return &Calendar{
		Name: name,
		Color: color,
		t: t,
	}
}

func NewWorkCalendar() *Calendar {
	c := NewCalendar("调休", "#DD2222")
	c.SetPrefix("🔴")
	return c
}

func NewOffCalendar() *Calendar {
	c := NewCalendar("节假日", "#22DD22")
	c.SetPrefix("🟢")
	return c
}


func (c *Calendar) SetPrefix(prefix string) {
	c.prefix = prefix
}

func (c *Calendar) Render(writer io.Writer) {
	c.t.ExecuteTemplate(writer, "icalendar.tmpl", c)
}

func (c *Calendar) RenderFile(path string) {
	f, _ := os.Create(path)
	defer f.Close()
	
	c.Render(f)
}

func (c *Calendar) RenderString() string {
	var b bytes.Buffer
	c.Render(&b)
	return b.String()
}

func (c *Calendar) AddHoliday(name string, date time.Time, rest bool) {
	c.Holidays = append(c.Holidays, Holiday{
		Name: c.prefix + name,
		StartDate: date.Format("20060102"),
		EndDate: date.AddDate(0, 0, 1).Format("20060102"),
		DTStamp: date.Format("20060102T150405Z"),
		Rest: rest,
	})
}
