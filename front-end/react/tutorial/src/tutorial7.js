/**
 * Created by Rosen on 2/13/17.
 */
var Comment = React.createClass({
  rawMarkup: function () {
    var md = new Remarkable();
    var rawMarkup = md.render(this.children.toString());
    return { __html: rawMarkup };
  },

  render: function () {
    return (
        <div className="comment">
          <h2 className="commentAuthor">
            {this.props.author}
          </h2>
          <span dangerouslySetInnerHTML={this.rawMarkup()} />
        </div>
    );
  }
});
