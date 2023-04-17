import React from 'react';
import LoginPage from './component/LoginPage';
import RegisterPage from './component/RegisterPage';
import './App.css';
function App() {
  return (
    <div className="contener">
      <div className="LoginPanel">
        <br />
        <h3>DianomiTV</h3>
        <LoginPage />
        <RegisterPage />
      </div>
    </div>
  );
}

export default App;
