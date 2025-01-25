import React, { useEffect, useState } from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import axios from "axios";

interface ThreadCardProps {
    categoryId: number;
    title: string;
    authorId: number;
    content: string;
    threadId: number;
    createdAt: string;
    buttonNeeded: boolean;
    isCreator: boolean;
}

const ThreadCard: React.FC<ThreadCardProps> = ({
    categoryId,
    title,
    authorId,
    content,
    threadId,
    createdAt,
    buttonNeeded,
    isCreator,
}) => {
    const [author, setAuthor] = useState<string | null>(null);
    const [category, setCategory] = useState<string | null>(null);

    isCreator;

    useEffect(() => {
        axios
            .get(`http://localhost:8080/api/v1/user/${authorId}`)
            .then((response) => {
                setAuthor(response.data);
            })
            .catch((error) => {
                console.error("There was an error fetching the username:", error);
            });
        axios
            .get(`http://localhost:8080/api/v1/categories/${categoryId}`)
            .then((response) => {
                setCategory(response.data);
            })
            .catch((error) => {
                console.error("There was an error fetching the category:", error);
            });
    }, [authorId, categoryId]);

    const currentUserId = localStorage.getItem("userId");

    const formatDate = (dateString: string): string => {
        const options: Intl.DateTimeFormatOptions = { month: "long", day: "numeric" };
        const date = new Date(dateString);
        return date.toLocaleDateString("en-US", options);
    };

    return (
        <Card
            variant="outlined"
            sx={{ minWidth: 550, marginBottom: 2, padding: 1, display: "flex", flexDirection: "column" }}
        >
            <CardContent sx={{ flexGrow: 1 }}>
                <Typography gutterBottom sx={{ color: "text.secondary", fontSize: 14, textAlign: "left" }}>
                    {category?.toUpperCase() || "Loading category..."}
                </Typography>
                <Typography variant="h5" component="div" sx={{ textAlign: "left", fontWeight: "bold" }}>
                    {title}
                </Typography>
                <Typography sx={{ color: "text.secondary", mb: 1.5, textAlign: "left" }}>{content}</Typography>
            </CardContent>
            <Box
                sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    paddingX: 2,
                    paddingBottom: 1,
                    alignItems: "center",
                }}
            >
                <Typography sx={{ color: "text.secondary", fontSize: 14 }}>
                    Posted by {author || "... (loading...)"} on {formatDate(createdAt)}
                </Typography>
                <CardActions
                    sx={{
                        display: "flex",
                        justifyContent: "flex-end",
                        paddingTop: 0,
                    }}
                >
                    {buttonNeeded && (
                        <Button variant="contained" href={`http://localhost:3000/threads/${threadId}`}>
                            VIEW
                        </Button>
                    )}

                    {/* Edit button: only show if the current user is the thread's author */}
                    {currentUserId && currentUserId === authorId.toString() && (
                        <Button
                            variant="contained"
                            sx={{ backgroundColor: "orange", "&:hover": { backgroundColor: "gold" } }}
                            href={`/threads/${threadId}/edit`} // Navigate to the edit page
                        >
                            EDIT
                        </Button>
                    )}

                    {/* Delete button: only show if the current user is the thread's author */}
                    {currentUserId && currentUserId === authorId.toString() && (
                        <Button
                            variant="contained"
                            color="error"
                            href={`http://localhost:3000/threads/${threadId}/delete`}
                        >
                            DELETE
                        </Button>
                    )}
                </CardActions>
            </Box>
        </Card>
    );
};

export default ThreadCard;
