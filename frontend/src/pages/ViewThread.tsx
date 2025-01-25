import Header from "../components/Header";
import ThreadCard from "../components/ThreadCard";
import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams, Link } from "react-router-dom";
import { Box, Typography, TextField, Button, Snackbar } from "@mui/material";

interface Comment {
    id: number;
    authorId: number;
    content: string;
    createdAt: string;
}

interface Thread {
    id: number;
    title: string;
    content: string;
    authorId: number;
    CategoryID: number;
    createdAt: string;
}

const ViewThread: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [thread, setThread] = useState<Thread | null>(null);
    const [comments, setComments] = useState<Comment[]>([]);
    const [commentAuthors, setCommentAuthors] = useState<Map<number, string>>(new Map());
    const [content, setContent] = useState("");
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const [openSnackbar, setOpenSnackbar] = useState(false);
    const [editingCommentId, setEditingCommentId] = useState<number | null>(null);
    const [editedContent, setEditedContent] = useState("");
    const [loggedInUserId, setLoggedInUserId] = useState<number | null>(null);

    useEffect(() => {
        console.log(localStorage);
        const userId = localStorage.getItem("userId");
        if (userId) {
            setLoggedInUserId(parseInt(userId, 10));
        }
    }, []);

    useEffect(() => {
        const fetchThreadAndComments = async () => {
            try {
                const threadResponse = await axios.get(`http://localhost:8080/api/v1/threads/${id}`);
                setThread(threadResponse.data);

                const commentsResponse = await axios.get(`http://localhost:8080/api/v1/threads/${id}/comments`);
                setComments(commentsResponse.data);

                const authorsMap = new Map<number, string>();
                for (const comment of commentsResponse.data) {
                    const authorResponse = await axios.get(`http://localhost:8080/api/v1/user/${comment.authorId}`);
                    authorsMap.set(comment.authorId, authorResponse.data);
                }
                setCommentAuthors(authorsMap);
            } catch (error) {
                console.error("Error fetching thread or comments:", error);
            }
        };
        if (id) {
            fetchThreadAndComments();
        }
    }, [id]);

    const handleAddComment = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const token = localStorage.getItem("jwt");
            if (!token) {
                setErrorMessage("You must be logged in to add a comment.");
                setOpenSnackbar(true);
                return;
            }

            await axios.post(
                `http://localhost:8080/api/v1/threads/${id}`,
                { content },
                { headers: { Authorization: `Bearer ${token}` } },
            );

            window.location.reload();
        } catch (error) {
            setErrorMessage("Failed to add comment. Please try again.");
            setOpenSnackbar(true);
        }
    };

    const handleDeleteComment = async (commentId: number) => {
        try {
            const token = localStorage.getItem("jwt");
            if (!token) {
                setErrorMessage("You must be logged in to delete a comment.");
                setOpenSnackbar(true);
                return;
            }

            await axios.delete(`http://localhost:8080/api/v1/threads/comment/${commentId}`, {
                headers: { Authorization: `Bearer ${token}` },
            });

            window.location.reload();
        } catch (error) {
            setErrorMessage("Failed to delete comment. Please try again.");
            setOpenSnackbar(true);
        }
    };

    const handleEditComment = (commentId: number, currentContent: string) => {
        setEditingCommentId(commentId);
        setEditedContent(currentContent);
    };

    const handleSaveEdit = async (commentId: number) => {
        try {
            const token = localStorage.getItem("jwt");
            if (!token) {
                setErrorMessage("You must be logged in to edit a comment.");
                setOpenSnackbar(true);
                return;
            }

            await axios.put(
                `http://localhost:8080/api/v1/threads/comment/${commentId}`,
                { content: editedContent },
                { headers: { Authorization: `Bearer ${token}` } },
            );

            window.location.reload();
        } catch (error) {
            setErrorMessage("Failed to edit comment. Please try again.");
            setOpenSnackbar(true);
        }
    };

    return (
        <>
            <Header />

            <Box sx={{ marginBottom: 0, display: "flex", justifyContent: "center", padding: 3 }}>
                <Link to="/">
                    <Button variant="outlined" color="primary">
                        Back to Home
                    </Button>
                </Link>
            </Box>

            {thread ? (
                <Box sx={{ padding: 3 }}>
                    <ThreadCard
                        categoryId={thread.CategoryID}
                        title={thread.title}
                        authorId={thread.authorId}
                        content={thread.content}
                        threadId={thread.id}
                        createdAt={thread.createdAt}
                        buttonNeeded={false}
                        isCreator={false}
                    />

                    <Typography variant="h6" sx={{ marginTop: 3, marginBottom: 3 }}>
                        Comments
                    </Typography>

                    {comments.length > 0 ? (
                        comments.map((comment) => (
                            <Box
                                key={comment.id}
                                sx={{
                                    marginBottom: 2,
                                    padding: 2,
                                    backgroundColor: "#f5f5f5",
                                    borderRadius: 2,
                                    position: "relative",
                                }}
                            >
                                <Typography variant="body2" sx={{ fontWeight: "bold", textAlign: "left" }}>
                                    {commentAuthors.get(comment.authorId) || "... (loading...)"}
                                </Typography>

                                {editingCommentId === comment.id ? (
                                    <>
                                        <TextField
                                            value={editedContent}
                                            onChange={(e) => setEditedContent(e.target.value)}
                                            fullWidth
                                            multiline
                                            rows={2}
                                            sx={{ marginTop: 1 }}
                                        />
                                        <Button
                                            onClick={() => handleSaveEdit(comment.id)}
                                            variant="contained"
                                            color="primary"
                                            fullWidth
                                            sx={{ marginTop: 1 }}
                                        >
                                            Save
                                        </Button>
                                    </>
                                ) : (
                                    <Typography variant="body2" sx={{ textAlign: "left", marginTop: 1 }}>
                                        {comment.content}
                                    </Typography>
                                )}

                                {loggedInUserId === comment.authorId && (
                                    <Box sx={{ position: "absolute", bottom: 8, right: 8, display: "flex", gap: 1 }}>
                                        <Button
                                            onClick={() => handleEditComment(comment.id, comment.content)}
                                            variant="outlined"
                                            color="primary"
                                        >
                                            Edit
                                        </Button>
                                        <Button
                                            onClick={() => handleDeleteComment(comment.id)}
                                            variant="outlined"
                                            color="error"
                                        >
                                            Delete
                                        </Button>
                                    </Box>
                                )}
                            </Box>
                        ))
                    ) : (
                        <Typography variant="body2" sx={{ marginBottom: 2 }}>
                            No comments yet.
                        </Typography>
                    )}

                    <form
                        onSubmit={handleAddComment}
                        style={{ display: "flex", flexDirection: "column", maxWidth: 2000, justifyContent: "right" }}
                    >
                        <Box sx={{ display: "flex", alignItems: "right", marginBottom: 2 }}>
                            <TextField
                                label="Add a Comment"
                                variant="outlined"
                                fullWidth
                                multiline
                                rows={2}
                                value={content}
                                onChange={(e) => setContent(e.target.value)}
                                sx={{ marginRight: 2 }}
                                required
                            />
                            <Button type="submit" variant="contained" color="primary" sx={{ height: "100%" }}>
                                Add Comment
                            </Button>
                        </Box>
                    </form>
                </Box>
            ) : (
                <Typography variant="h6" sx={{ padding: 3 }}>
                    Loading thread...
                </Typography>
            )}

            <Snackbar
                open={openSnackbar}
                autoHideDuration={6000}
                onClose={() => setOpenSnackbar(false)}
                message={errorMessage}
            />
        </>
    );
};

export default ViewThread;
