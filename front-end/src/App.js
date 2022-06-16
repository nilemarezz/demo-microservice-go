import {
  BrowserRouter as Router,
  Routes,
  Route,
  useRoutes,
  Navigate,
} from "react-router-dom";
import MovieDetail from "./container/MovieDetail";
import MovieList from "./container/MovieList";
import { useEffect, useState } from "react";
import LoadingScreen from "./components/LoadingScreen";
import Login from "./container/Login";
import Register from "./container/Register";
import axios from "axios";

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  return (
    <div className="App">
      <Router>
        <RouteList isLoggedIn={isLoggedIn} setIsLoggedIn={setIsLoggedIn} />
      </Router>
    </div>
  );
};

const RouteList = ({ isLoggedIn, setIsLoggedIn, movies }) => {
  let routes = useRoutes([
    {
      path: "/movies",
      element: isLoggedIn ? <MovieList /> : <Navigate to="/login" />,
    },
    {
      path: "/movies/:id",
      element: isLoggedIn ? <MovieDetail /> : <Navigate to="/login" />,
    },
    { path: "/login", element: <Login setIsLoggedIn={setIsLoggedIn} /> },
    { path: "/register", element: <Register /> },
    { path: "/", element: <Navigate to="/movies" /> },
  ]);
  return routes;
};

export default App;
