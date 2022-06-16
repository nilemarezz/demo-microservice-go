import { Button, Typography } from "@mui/material";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import axios from "axios";

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  let navigate = useNavigate();
  const onSubmit = async () => {
    try {
      const data = await axios.post("http://127.0.0.1:5000/auth/signup", {
        username,
        password,
      });
      console.log(data);
      if (data.data.success) {
        navigate("/login");
      } else {
        setError("Signup fail");
      }
    } catch (err) {
      console.log(err);
      setError("Signup fail");
    }
  };
  return (
    <div>
      <Box
        component="form"
        sx={{
          "& > :not(style)": { m: 1, width: "25ch" },
        }}
        noValidate
        autoComplete="off"
      >
        <Typography variant="h5" component="h5">
          Register
        </Typography>

        <TextField
          id="outlined-basic"
          label="Username"
          variant="outlined"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <TextField
          id="outlined-basic"
          label="Password"
          type={"password"}
          variant="outlined"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div>
          <Button variant="contained" onClick={onSubmit}>
            Register
          </Button>
          <Link to="/login">
            <Typography>Have account?</Typography>
          </Link>
        </div>
        {error && <Typography>{error}</Typography>}
      </Box>
    </div>
  );
};

export default Register;
