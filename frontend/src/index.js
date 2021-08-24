import * as React from 'react'
import * as ReactDOM from 'react-dom'
import * as Wails from '@wailsapp/runtime'
import {App} from './App'
import 'core-js/stable'
import './index.css'

Wails.Init(() => {
  ReactDOM.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
    document.getElementById('app')
  )
})
