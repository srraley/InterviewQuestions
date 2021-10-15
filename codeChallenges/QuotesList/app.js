/** Your working area*/
class ClassView extends React.Component{
    render(){
        return(<h1>Quotes Challenge</h1>)
    }
}

const FunctionView = function(){
    // you can access hooks by Refering to the React object, for example const [state,setState] = React.useState
    return(<h1>Quotes Challenge</h1>)
}


/** Not working area*/
const App = React.createElement(ClassView);
ReactDOM.render(
    App,
    document.getElementById('root')
  );