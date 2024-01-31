package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Comment struct {
	ID       int
	ParentID int
	Text     string
}

type Comments []*Comment
type CommentsMap map[int][]*Comment

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var setCounter int
	fmt.Fscanln(in, &setCounter)

	for i := 0; i < setCounter; i++ {

		var commentCounter int
		fmt.Fscanln(in, &commentCounter)

		comments := make(Comments, 0, commentCounter)

		for j := 0; j < commentCounter; j++ {

			var id int
			var parentID int
			var text string
			_, err := fmt.Fscan(in, &id)
			if err != nil {
				panic(err)
			}

			_, err = fmt.Fscan(in, &parentID)
			if err != nil {
				panic(err)
			}

			text, err = in.ReadString('\n')
			if err != nil {
				panic(err)
			}

			comments = append(comments, &Comment{
				ID:       id,
				ParentID: parentID,
				Text:     strings.Trim(text[1:], "\n"),
			})

		}

		sort.SliceStable(comments, func(i, j int) bool { return comments[i].ID < comments[j].ID })

		commentsMap := CommentsMap{}

		for _, comment := range comments {
			commentsMap[comment.ParentID] = append(commentsMap[comment.ParentID], comment)
		}

		tree := buildTree(commentsMap)
		fmt.Fprint(out, tree)
		if i < setCounter-1 {
			fmt.Fprint(out, "\n")
		}

	}

}

func buildTree(comments CommentsMap) string {

	sb := &strings.Builder{}

	next := map[int]bool{}
	buildLevel(sb, comments, -1, 1, next)

	return sb.String()

}

func buildLevel(sb *strings.Builder, comments CommentsMap, parentID int, level int, next map[int]bool) {

	levelComments := comments[parentID]
	if len(levelComments) == 0 {
		return
	}

	for i, comment := range levelComments {

		if level == 1 {
			if i > 0 {
				sb.WriteString("\n")
			}
			sb.WriteString(comment.Text)
			sb.WriteString("\n")
		} else {
			for i := 2; i <= level; i++ {
				if i == level {
					sb.WriteString("|")
				} else {
					nextCommentsExist := next[i]
					if nextCommentsExist {
						sb.WriteString("|  ")
					} else {
						sb.WriteString("   ")
					}
				}
			}

			sb.WriteString("\n")

			for i := 2; i <= level; i++ {
				if i == level {
					sb.WriteString("|")
				} else {
					nextCommentsExist := next[i]
					if nextCommentsExist {
						sb.WriteString("|  ")
					} else {
						sb.WriteString("   ")
					}
				}
			}

			sb.WriteString("--")
			sb.WriteString(comment.Text)
			sb.WriteString("\n")

		}

		next[level] = false
		if i < len(levelComments)-1 {
			next[level] = true
		}

		buildLevel(sb, comments, comment.ID, level+1, next)

	}

}

func str2int(s string) int {
	res, err := strconv.Atoi(strings.Trim(s, " "))
	if err != nil {
		panic(err)
	}
	return res
}
