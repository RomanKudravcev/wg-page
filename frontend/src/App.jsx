import { Route, Routes, Navigate } from "react-router-dom";
import Dashboard from "./Pages/Dashboard";
import Todo from "./Pages/Todo";
import Login from "./Pages/Login"
import "./App.css";

function App() {
    return (
        <div className="App">
            <Routes>
                <Route path="/dashboard" element={<Dashboard />}></Route>
                <Route path="/todo" element={<Todo />}></Route>
                <Route path="/login" element={<Login />}></Route>
                <Route path="*" element={<Navigate to="/login" />}></Route>
            </Routes>
        </div >
    );
}

export default App;
