/**
 * Created by Rosen on 2/13/17.
 */
var CommentBox = React.createClass({
  render: function () {
    return (
        <div className="commentBox">
          <h1>Comments</h1>
          <CommentList />
          <CommentForm />
        </div>
    );
  }
});
