import Home from "./pages/Home";
import "./App.css";
import Login from "./pages/Login";
import Register from "./pages/Register";
import CreateThread from "./pages/CreateThread";
import ViewThread from "./pages/ViewThread";
import DeleteThread from "./pages/DeleteThread";
import EditThread from "./pages/EditThread";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { blue, orange } from "@mui/material/colors";
import React from "react";

const theme = createTheme({
    palette: {
        primary: blue,
        secondary: orange,
    },
});

const App: React.FC = () => {
    return (
        <div className="App">
            <ThemeProvider theme={theme}>
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/login" element={<Login />} />
                        <Route path="/register" element={<Register />} />
                        <Route path="/threads/new" element={<CreateThread />} />
                        <Route path="/threads/:id" element={<ViewThread />} />
                        <Route path="/threads/:id/delete" element={<DeleteThread />} />
                        <Route path="/threads/:id/edit" element={<EditThread />} />
                    </Routes>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default App;
