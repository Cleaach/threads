import React, { useEffect, useState } from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

const Header: React.FC = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    useEffect(() => {
        const fetchUserInfo = async () => {
            try {
                const token = localStorage.getItem("jwt");
                if (!token) {
                    setIsLoggedIn(false);
                    return;
                }
                setIsLoggedIn(true);
            } catch (error) {
                console.error("Error fetching user info:", error);
                setIsLoggedIn(false);
            }
        };
        fetchUserInfo();
    }, []);

    const handleLogout = () => {
        localStorage.removeItem("jwt");
        setIsLoggedIn(false);
    };

    return (
        <AppBar position="static" color="primary">
            <Toolbar
                sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                }}
            >
                <Box
                    sx={{
                        width: 400,
                        display: "flex",
                        justifyContent: "flex-start",
                    }}
                ></Box>
                <Typography
                    variant="h6"
                    sx={{
                        textAlign: "center",
                    }}
                >
                    Threads
                </Typography>
                <Box
                    sx={{
                        width: 400,
                        display: "flex",
                        justifyContent: "flex-end",
                        gap: 1,
                    }}
                >
                    {isLoggedIn ? (
                        <>
                            <Button color="inherit" onClick={handleLogout}>
                                Log Out
                            </Button>
                        </>
                    ) : (
                        <>
                            <Button color="inherit" href="/register">
                                Sign Up
                            </Button>
                            <Button color="inherit" href="/login">
                                Log In
                            </Button>
                        </>
                    )}
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default Header;
