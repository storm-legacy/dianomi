import React, { FormEvent, useEffect, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import authService from '../../services/auth.service';

export const ResetPage = () => {
  interface ResetData {
    validate: string;
    password: string;
    password_repeat: string;
  }
  const [data, setData] = useState<ResetData>();
  const [isError, setIsError] = useState(false);
  const [isSuccess, setIsSuccess] = useState(false);
  const [Password, setPassword] = useState('');
  const [PasswordRepeat, setPasswordRepeat] = useState('');
  const [isValid, setIsValid] = useState(true);
  const [validate, setValidate] = useState('');
  const navigate = useNavigate();
  useEffect(() => {
    const searchParams = new URLSearchParams(location.search);
    const validate = searchParams.get('validate');
    const validateCode = validate || '';
    console.log(validate);
    setValidate(validateCode);
    const { request } = authService.resetGet(validateCode);
    request
      .then((res) => {
        console.log(res);
        setIsError(false);
      })
      .catch((err) => {
        console.error(err);
        setIsError(true);
      });
  }, []);
  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    const { request } = authService.resetPost(validate, Password, PasswordRepeat);
    request
      .then((res) => {
        console.log(res);
        setIsSuccess(true);
      })
      .catch((err) => {
        console.error(err.message);
        setIsSuccess(false);
      });
  };

  if (isError == true) {
    return (
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded ">
        <div className="alert alert-danger">
          Something went wrong please go back <Link to={'/login'}>Login page</Link>{' '}
        </div>
      </div>
    );
  }
  return (
    <div className="MyRes position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
      <h3>Reset your password</h3>
      <form onSubmit={handleSubmit}>
        <br />
        <label>
          <p className="h5">Password:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            aria-label="Password"
            aria-describedby="basic-addon1"
            value={Password}
            onChange={(event) => setPassword(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Repeat password:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Repeat password"
            aria-label="PasswordRepeat"
            aria-describedby="basic-addon1"
            value={PasswordRepeat}
            onChange={(event) => setPasswordRepeat(event.target.value)}
          />
        </label>

        <br />
        <button type="submit" className="btn btn-primary">
          Create
        </button>
      </form>
      {!isValid && (
        <p className="alert alert-danger" style={{ marginTop: '5px' }}>
          Please enter a valid email.
        </p>
      )}
      {isSuccess && (
        <p className="alert alert-success " style={{ marginTop: '5px' }}>
          The operation was successful, please go back to login panel <Link to={'/login'}>Login page</Link>
        </p>
      )}
      {Password != PasswordRepeat && PasswordRepeat != '' && (
        <p className="alert alert-danger" style={{ marginTop: '5px' }}>
          {' '}
          The passwords are not identical <br />{' '}
        </p>
      )}
    </div>
  );
};
