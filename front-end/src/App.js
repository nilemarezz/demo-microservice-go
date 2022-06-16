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

const App = () => {
  const [movies, setMovies] = useState(null);
  const [loading, setLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  useEffect(() => {
    setLoading(true);

    const getMovies = async () => {
      setLoading(true);
      const res = await fetch("http://127.0.0.1:5000/movies/", {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      const data = await res.json();
      if (data.error) {
        setLoading(false);
      } else {
        setMovies(data);
        setIsLoggedIn(true);
        setLoading(false);
      }
    };

    getMovies();
  }, []);

  if (loading) {
    return <LoadingScreen />;
  }
  return (
    <div className="App">
      <Router>
        <RouteList
          isLoggedIn={isLoggedIn}
          setIsLoggedIn={setIsLoggedIn}
          movies={movies}
        />
      </Router>
    </div>
  );
};

const RouteList = ({ isLoggedIn, setIsLoggedIn, movies }) => {
  let routes = useRoutes([
    {
      path: "/movies",
      element: isLoggedIn ? (
        <MovieList movies={movies} />
      ) : (
        <Navigate to="/login" />
      ),
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
