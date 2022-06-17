import { Paper, Typography, Grid, Button } from "@mui/material";
import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import LoadingScreen from "../components/LoadingScreen";
import axios from "axios";

const MovieList = () => {
  const [loading, setLoading] = useState(false);
  const [movies, setMovies] = useState(null);

  useEffect(() => {
    const getMovies = async () => {
      setLoading(true);
      try {
        const res = await fetch("http://127.0.0.1:5000/movies/", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + localStorage.getItem("token"),
          },
        });
        const data = await res.json();
        console.log(data);
        setMovies(data);
        setLoading(false);
      } catch (err) {
        console.log(err);
        setLoading(false);
      }
    };
    getMovies();
  }, []);

  if (loading) {
    return <LoadingScreen />;
  }

  return (
    <div>
      <Typography variant="h5" component="h5">
        Movie List
      </Typography>
      <Grid container spacing={2}>
        {movies &&
          movies.map((movie) => (
            <MovieItem
              key={movie.id}
              id={movie.id}
              name={movie.name}
              screenDate={movie.screen_date}
            />
          ))}
      </Grid>
    </div>
  );
};

const MovieItem = ({ id, name, screenDate }) => {
  return (
    <Grid item xs={12} sm={6} md={3}>
      <Paper elevation={3}>
        <div style={{ padding: 3 }}>
          <Typography variant="h6" component="h6">
            {name}
          </Typography>
          <Typography>{screenDate}</Typography>
          <Link to={`/movie/${id}`} style={{ textDecoration: "none" }}>
            <Button variant="contained" fullWidth>
              Detail
            </Button>
          </Link>
        </div>
      </Paper>
    </Grid>
  );
};

export default MovieList;
