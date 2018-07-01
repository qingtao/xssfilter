package xssfilter

var (
	before = `
<div>
<h1>这是文本标题</h1>
</div>
<div>
<p>xss</p>
<script>console.log('xss testing')</script>
</div>
<div>
<script>console.log('xss testing')</script>
<p>这是段落前
<span>abc</span>
</p>
</div>
<div>
<p>这是文本内容1,正常</p>
<p onclick="console.log('hello')">这是文本内容2,包含onclick</p>
<p class="text3" id="text3" style="color:red">这是文本内容3,包含class/id/style</p>
<p onmouseup="console.log('hello2')">这是文本内容4,包含onmouseup/script
<script>console.log('haha')</script>
</p>
</div>`
	after = `<div>
<h1>这是文本标题</h1>
</div>
<div>
<p>xss</p>

</div>
<div>

<p>这是段落前
<span>abc</span>
</p>
</div>
<div>
<p>这是文本内容1,正常</p>
<p>这是文本内容2,包含onclick</p>
<p class="text3" id="text3" style="color:red">这是文本内容3,包含class/id/style</p>
<p>这是文本内容4,包含onmouseup/script

</p>
</div>`
)
