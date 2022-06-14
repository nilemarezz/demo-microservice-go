import { Paper, Typography, Grid, Button } from "@mui/material";
import { Link } from "react-router-dom";

const MovieList = ({ movies }) => {
  return (
    <div>
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
