import React, { useEffect, useState } from "react";
import { axiosInstance } from "../utils/axios";
import { useNavigate } from "react-router-dom";
import { Button, Container, Stack, TextField, Typography, Paper, Box } from "@mui/material";

const ProfilePage = () => {
  const [profile, setProfile] = useState(null);
  const navigate = useNavigate();
  const [title, setTitle] = useState('')
  const [todos, setTodos] = useState([])
  useEffect(() => {
    init();
  }, []);
  const init = async () => {
    const response = await axiosInstance.get("/api/profile");
    if (response?.data?.status !== "200") {
      return navigate("/sign-in");
    }
    console.log("🚀 ===== init ===== response:", response);
    setProfile(response?.data?.data);
    refetch()
  };
  const handleChangeTitle = (e) => {
    const value = e.target.value
    setTitle(value)
  }
  const handleCreate = async () => {
    const form = new FormData()
    form.append('title',title)
    const response = await axiosInstance.post("/api/todo/create",form);
    if(response?.status===200 && response?.data?.status==="200"){
      refetch()
    }
  };
  const refetch = async () => {
    const response = await axiosInstance.get("/api/todos");
    if(response?.status===200 && response?.data?.status==="200"){
      setTodos(response?.data?.data)
    }
  };
  return (
    <Container maxWidth="sm">
      <Stack direction="row" alignItems="center" justifyContent="space-between">
        <h2>Todolist</h2>
        <Stack direction="row" alignItems="center">
          <p>Hello {profile?.Email}</p>
          <Button size="small" sx={{height:'fit-content'}} variant="contained">
            Logout
          </Button>
        </Stack>
      </Stack>
      <Stack sx={{ width: "100%" }} direction="row" alignItems="center" spacing={2}>
        <TextField fullWidth={true} onChange={handleChangeTitle}></TextField>
        <Button variant="contained" sx={{height:'100%'}} onClick={handleCreate}>Add</Button>
      </Stack>
      <Todolist todos={todos}/>
    </Container>
  );
};
const Todolist = ({todos}) => {
  return (
    <Stack direction="column" spacing={2} sx={{mt:3}}>
      {todos?.map((todo) => <Todo todo={todo} />)}
    </Stack>
  )
}
const Todo = ({todo})=>{
  return (
    <Paper sx={{p:2}}>
      <Stack >
        <Typography>{todo?.Title}</Typography>
      </Stack>
    </Paper>
  )
}

export default ProfilePage;
