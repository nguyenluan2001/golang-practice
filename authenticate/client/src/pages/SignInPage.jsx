import React from "react";
import {
  Container,
  Typography,
  Stack,
  Box,
  TextField,
  Button,
  Paper,
} from "@mui/material";
import { Controller, useForm } from "react-hook-form";
import { axiosInstance } from "../utils/axios";
import axios from "axios";
import {useNavigate} from "react-router-dom"

const SignInPage = () => {
  const navigate = useNavigate()
  const { control, handleSubmit } = useForm({
    defaultValues: {
      email: "",
      password: "",
    },
  });
  const onSubmit = async (values) => {
    console.log("🚀 ===== onSubmit ===== values:", values);
    const { email, password } = values;
    const form = new FormData();
    form.append("email", email);
    form.append("password", password);
    const response = await axiosInstance.post('http://localhost:8081/api/sign-in',form)
    if(response?.status===200 && response?.data?.status==="200"){
      navigate('/')
    }
  };
  return (
    <Container fullWidth={true} maxWidth="xs">
      <Paper
        disableGutters={true}
        sx={{ p: 2, minHeight: "50vh", mt: "100px" }}
        elevation={3}
      >
        <Stack direction="column" alignItems="center" spacing={2}>
          <Typography variant="h3">Sign in</Typography>
          <Stack sx={{ width: "100%" }} direction="column" spacing={2}>
            <Stack
              sx={{ width: "100%" }}
              direction="column"
              alignItems="flex-start"
            >
              <Typography>Email</Typography>
              <Controller
                control={control}
                name="email"
                render={({ field }) => (
                  <TextField fullWidth={true} {...field} />
                )}
              ></Controller>
            </Stack>
            <Stack
              sx={{ width: "100%" }}
              direction="column"
              alignItems="flex-start"
            >
              <Typography>Password</Typography>
              <Controller
                control={control}
                name="password"
                render={({ field }) => (
                  <TextField fullWidth={true} type="password" {...field} />
                )}
              ></Controller>
            </Stack>
          </Stack>
          <Button variant="contained" onClick={handleSubmit(onSubmit)}>
            Sign in
          </Button>
        </Stack>
      </Paper>
    </Container>
  );
};

export default SignInPage;
