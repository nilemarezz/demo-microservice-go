import { BrowserRouter, Routes, Route } from "react-router-dom";
import MovieDetail from "./container/MovieDetail";
import MovieList from "./container/MovieList";
import { useEffect, useState } from "react";
import LoadingScreen from "./components/LoadingScreen";

const App = () => {
  const [movies, setMovies] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const getMovies = async () => {
      setLoading(true);
      const res = await fetch("http://127.0.0.1:5000/movies/");
      const data = await res.json();
      setMovies(data);
      setLoading(false);
    };

    getMovies();
  }, []);

  if (loading) {
    return <LoadingScreen />;
  }
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<MovieList movies={movies} />}></Route>
          <Route path="/movie/:id" element={<MovieDetail />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
};

export default App;
