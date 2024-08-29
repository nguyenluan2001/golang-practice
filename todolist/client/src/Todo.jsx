import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {
  useQuery,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'


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

function Todo() {
    const {data} = useQuery({
        queryKey: ['fetch-todos'],
        queryFn: () => fetch('http://localhost:8081/api/todos').then((res) => res.json())
    })
    console.log('data', data)
  return (
    <div className="flex flex-col items-center justify-center w-[60%] mx-auto">
      <h1 className="h-10">Todolist Golang</h1>
      <div className="flex w-full mb-3 gap-2">
        <input type="text" className="flex-grow outline"></input>
        <button className="bg-slate-500 text-white rounded-sm min-w-[50px] py-1 px-2">Create</button>
      </div>
      <div className="w-full flex flex-col gap-1">
        {
          data?.map((todo) => {
            return (
              <div className="w-full p-[8px] border flex gap-2" key={todo.uid}>
                <p>{todo.title}</p>
                <div className={`${todo.status==='DONE' ? 'bg-emerald-600 text-white' : 'bg-yellow-300'} text-[12px] border-rounded font-medium py-1 px-2`}>{todo.status}</div>
              </div>
            )
          })
        }

      </div>
    </div>
  )
}

export default Todo