import React, { useState } from "react";
import { Button, Typography, Box } from "@mui/material";
import axios from "axios";
import { useNavigate, useParams } from "react-router-dom";

const DeleteThread: React.FC = () => {
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const navigate = useNavigate();
    const { id } = useParams<{ id: string }>();

    const handleDelete = async () => {
        try {
            const token = localStorage.getItem("jwt");

            if (!token) {
                setErrorMessage("You must be logged in to delete this thread.");
                return;
            }
            const response = await axios.delete(`http://localhost:8080/api/v1/threads/${id}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            console.log(response);
            navigate("/");
        } catch (error) {
            setErrorMessage("Failed to delete the thread. Please try again.");
            console.error("Error deleting thread:", error);
        }
    };

    return (
        <Box sx={{ padding: 3 }}>
            <Typography variant="h6" sx={{ marginBottom: 2 }}>
                Are you sure you want to delete this thread?
            </Typography>
            {errorMessage && (
                <Typography color="error" sx={{ marginBottom: 2 }}>
                    {errorMessage}
                </Typography>
            )}
            <Button variant="contained" color="error" onClick={handleDelete}>
                Delete Thread
            </Button>
        </Box>
    );
};

export default DeleteThread;
