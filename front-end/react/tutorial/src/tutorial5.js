/**
 * Created by Rosen on 2/13/17.
 */
var CommentList = React.createClass({
  render: function () {
    return (
        <div className="commentList">
          <Comment author="Pete Hunt">This is one comment</Comment>
          <Comment author="Jordan Walke">This is *another* comment</Comment>
        </div>
    );
  }
});
