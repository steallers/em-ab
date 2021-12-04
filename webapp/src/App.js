import './App.css';
import {BrowserRouter, Link, Route, Routes} from "react-router-dom";


function App() {
    return (
        // Header
        // Some Menu
        <BrowserRouter>
            <Routes>
                <Route exact path="/" element={<Home/>}/>
                {/*<Route path="/about" component={About} />*/}
                {/*<Route path="/contact" component={Contact} />*/}

            </Routes>
        </BrowserRouter>
    );
}

function Home() {
    return (
        <div>
            <h1>Home</h1>
            <Link to="/about">About</Link>
            <Link to="/contact">Contact</Link>
        </div>
    );
}

export default App;
