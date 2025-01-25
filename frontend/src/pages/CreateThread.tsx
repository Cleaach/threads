import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import {
    TextField,
    Button,
    Container,
    Box,
    Typography,
    Snackbar,
    MenuItem,
    Select,
    InputLabel,
    FormControl,
} from "@mui/material";
import axios from "axios";

const CreateThread: React.FC = () => {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const [category, setCategory] = useState("");
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const [openSnackbar, setOpenSnackbar] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const token = localStorage.getItem("jwt");
            if (!token) {
                setErrorMessage("You must be logged in to create a thread.");
                setOpenSnackbar(true);
                return;
            }
            const response = await axios.post(
                `http://localhost:8080/api/v1/threads`,
                { title, content, category },
                {
                    headers: { Authorization: `Bearer ${token}` },
                },
            );
            response;
            navigate(`/`);
        } catch (error) {
            setErrorMessage("Failed to create thread. Please try again.");
            setOpenSnackbar(true);
        }
    };

    return (
        <Container>
            <Box
                sx={{
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "center",
                    alignItems: "center",
                    paddingTop: 2,
                }}
            >
                <Typography variant="h4" gutterBottom>
                    Create a New Thread
                </Typography>

                <form onSubmit={handleSubmit} style={{ width: "100%", maxWidth: 600 }}>
                    <TextField
                        label="Title"
                        variant="outlined"
                        fullWidth
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        sx={{ marginBottom: 2 }}
                        required
                    />
                    <TextField
                        label="Content"
                        variant="outlined"
                        fullWidth
                        multiline
                        rows={6}
                        value={content}
                        onChange={(e) => setContent(e.target.value)}
                        sx={{ marginBottom: 2 }}
                        required
                    />

                    <FormControl fullWidth sx={{ marginBottom: 2 }} required>
                        <InputLabel>Category</InputLabel>
                        <Select value={category} onChange={(e) => setCategory(e.target.value)} label="Category">
                            <MenuItem value="General Discussion">General Discussion</MenuItem>
                            <MenuItem value="Technology">Technology</MenuItem>
                            <MenuItem value="Science">Science</MenuItem>
                            <MenuItem value="Lifestyle">Lifestyle</MenuItem>
                        </Select>
                    </FormControl>

                    <Button type="submit" variant="contained" color="primary" sx={{ width: "100%" }}>
                        Create Thread
                    </Button>
                </form>
                <Snackbar
                    open={openSnackbar}
                    autoHideDuration={6000}
                    onClose={() => setOpenSnackbar(false)}
                    message={errorMessage}
                />
            </Box>
        </Container>
    );
};

export default CreateThread;
