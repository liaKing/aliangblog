import MarkdownIt from 'markdown-it';
export default function markd(data) {
    var md = new MarkdownIt();
    var result = md.render(data);
    return result
}
