This is a simple plugin for the static site generator [Booklit](https://vito.github.io/booklit/) that adds support for footnotes. As in LaTeX, use `\footnote` to define and `\footnotemark` to display them. Each call of `\footnote` will add its argument to a global list of notes (which is a bad idea). `\footnotemark` prints and empties the list.

To use it, copy both files inside `html` into your theme folder and add the plugin to the plugin list:

```bash
$ booklit \
    --plugin github.com/BakeRolls/booklit-footnotes \
    --html-templates html/ \
    # ...
```

```lit
\use-plugin{footnote}

\title{Footnotes example}

Just an example\footnote{An exampleUsing footnotes.}.

\footnotemark
```

Will render as

<div style="margin: 1em 2em;">
	<h1 class="section-header">
		<a name="footnotes-example"></a>
		<span class="section-number">2</span>
		Footnotes example
	</h1>
	<p>Just an example<sup><a href="#fn-0" name="fnref-0">0</a></sup>.</p>
	<p>
		<ol>
			<li>
				<a name="fn-0"></a>
				An exampleUsing footnotes.
				<sup><a href="#fnref-0">↩</a></sup>
			</li>
		</ol>
	</p>
</div>

This plugin should probably not be used.
