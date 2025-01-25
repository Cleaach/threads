import React, { useState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Login: React.FC = () => {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const response = await axios.post("http://localhost:8080/api/v1/login", {
                username,
                password,
            });

            if (response.status === 200) {
                localStorage.setItem("jwt", response.data.token);
                localStorage.setItem("userId", response.data.userId);
                navigate("/");
            }
        } catch (error) {
            setErrorMessage("Invalid credentials. Please try again.");
        }
    };

    return (
        <Box sx={{ width: "30vw", margin: "auto", paddingTop: 5 }}>
            <Typography variant="h4" gutterBottom>
                Log In
            </Typography>
            <form onSubmit={handleSubmit}>
                <TextField
                    label="Username"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <TextField
                    label="Password"
                    type="password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                {errorMessage && (
                    <Typography color="error" variant="body2">
                        {errorMessage}
                    </Typography>
                )}
                <Button type="submit" variant="contained" fullWidth sx={{ marginTop: 2 }}>
                    Log In
                </Button>
            </form>
        </Box>
    );
};

export default Login;
