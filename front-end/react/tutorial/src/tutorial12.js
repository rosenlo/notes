/**
 * Created by Rosen on 2/14/17.
 */
var CommentBox = React.createClass({
  getInitialState: function () {
    return {data: []};
  },
  render: function () {
    return (
        <div className="commentBox">
          <h1>Comments</h1>
          <CommentList data={this.state.data} />
          <CommentForm />
        </div>
    );
  }
});
