const App = React.createClass({
  render() {
    return (
        <div>
          <h1>App</h1>
          <ul>
            <li><Link to="/about">About</Link></li>
            <li><Link to="/index">Index</Link></li>
          </ul>
          {this.props.children}
        </div>
    )
  }
});

const About = React.createClass({
  render() {
    return <h3>About</h3>
  }
});

const Index = React.createClass({
  render() {
    return (
        <div>
          <h2>Inbox</h2>
          {this.props.chilren || "Welcome to your Inbox"}
        </div>
    )
  }
});

const Message = React.createClass({
  render() {
    return <h3>Message {this.props.params.id}</h3>
  }
});

React.render((
    <Router>
      <Route path="/" component={App}>
        <Route path="about" component={About}/>
        <Route path="index" component={Index}>
          <Route path="messages/:id" component={Message}/>
        </Route>
      </Route>
    </Router>
), document.getElementById('App'));