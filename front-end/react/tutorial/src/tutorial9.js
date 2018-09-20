/**
 * Created by Rosen on 2/14/17.
 */
var CommentBox = React.createClass({
  render: function () {
    return (
        <div className="commentBox">
          <h1>Comments</h1>
          <CommentList data={this.props.data}/>
          <CommentForm />
        </div>
    );
  }
});

ReactDOM.render(
    <CommentBox data={data}/>,
    document.getElementById('content')
);
