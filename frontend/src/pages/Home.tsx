import BasicThreadList from "../components/BasicThreadList";
import Header from "../components/Header";
import React from "react";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

const Home: React.FC = () => {
    return (
        <>
            <Header />
            <Box sx={{ display: "flex", justifyContent: "center", marginBottom: 2, marginTop: 5 }}>
                <Button
                    href="http://localhost:3000/threads/new"
                    variant="contained"
                    color="primary"
                    sx={{ width: 150 }}
                >
                    POST A THREAD
                </Button>
            </Box>
            <BasicThreadList />
        </>
    );
};

export default Home;
