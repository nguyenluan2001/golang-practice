import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {
  useQuery,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import Todo from './Todo'
const queryClient = new QueryClient()


const data = [
  {
    uid: 1,
    title: 'Todo 1',
    status: 'DONE',
  },
  {
    uid: 2,
    title: 'Todo 2',
    status: 'TODO',
  },
  {
    uid: 3,
    title: 'Todo 3',
    status: 'TODO',
  }
]

function App() {
  return (
    <QueryClientProvider client={queryClient}>
        <Todo />
    </QueryClientProvider>
  )
}

export default App
