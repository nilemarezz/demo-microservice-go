import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Typography } from "@mui/material";
import LoadingScreen from "../components/LoadingScreen";

const MovieDetail = () => {
  let { id } = useParams();

  const [detail, setDetail] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const getMovieDetail = async () => {
      setLoading(true);
      try {
        const res = await fetch(`http://127.0.0.1:5000/movies/${id}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + localStorage.getItem("token"),
          },
        });
        const data = await res.json();
        console.log(data);
        setDetail(data);
        setLoading(false);
      } catch (err) {
        console.log(err);
        setLoading(false);
      }
    };
    getMovieDetail();
  }, []);

  if (loading) {
    return <LoadingScreen />;
  }

  if (detail) {
    return (
      <div>
        <Typography variant="h5" component="h5">
          {detail.name}
        </Typography>
        <Typography>{detail.description}</Typography>
        <br></br>
        <Typography>Date : {detail.screen_date}</Typography>
        <Typography>Cast : </Typography>
        {detail &&
          detail.cast.map((item) => (
            <Typography key={item.id}>{item.name}</Typography>
          ))}
      </div>
    );
  }
};

export default MovieDetail;
