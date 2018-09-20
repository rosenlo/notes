/**
 * Created by Rosen on 2/13/17.
 */
var Comment = React.createClass({
  render: function () {
    return (
        <div className="comment">
          <h2 className="commentAuthor">
            {this.props.children}
          </h2>>
        </div>
    );
  }
});
