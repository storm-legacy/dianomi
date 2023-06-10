import React, { useContext, useEffect, useState } from 'react';
import { FiUser } from 'react-icons/fi';
import { AuthContext } from '../../context/AuthContext';
import profileService, { emailData, userResData, oldPass } from '../../services/profile.service';
import { Notify } from 'notiflix';
import { Package } from '../../services/admin.service';
import { CanceledError } from 'axios';
import { redirect, useNavigate } from 'react-router-dom';

export const ProfilePage = () => {
  const { user } = useContext(AuthContext);
  const navigate = useNavigate();
  const [isValiRep, setIsValiRep] = useState<string>('');
  const [isSame, setIsSame] = useState<boolean>();
  const [isValiOld, setIsValiOld] = useState<string>('');
  const [isValiNew, setIsValiNew] = useState<string>('');
  const [isError, setIsError] = useState<boolean>(true);
  const [Password, setPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [PasswordRepeat, setPasswordRepeat] = useState('');
  const [pack, setPack] = useState<Package>();
  const [daysLeft, setDaysLeft] = useState<number | null>(null);

  useEffect(() => {
    const { request, cancel } = profileService.GetPackage();
    request
      .then((res) => {
        setPack(res.data.data as Package);
        calculateDaysLeft(res.data.data);
      })
      .catch((err) => {
        if (err instanceof CanceledError) return;

        Notify.failure('Could not return information about self');
        console.error(err.message);
      });

    return () => cancel();
  }, []);

  const calculateDaysLeft = (pack: Package | undefined) => {
    if (pack?.valid_from && pack?.valid_until) {
      const validFromDate = new Date(pack.valid_from);
      const validUntilDate = new Date(pack.valid_until);
      const differenceInTime = validUntilDate.getTime() - validFromDate.getTime();
      const differenceInDays = Math.ceil(differenceInTime / (1000 * 3600 * 24));
      setDaysLeft(differenceInDays);
    }
  };
  const checkPasswordStrength = (password: string) => {
    const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,}$/;
    return regex.test(password) ? 1 : 0;
  };
  const handleSubscribe = () => {
    const { request } = profileService.PostPayment();
    request
      .then(({ data }) => {
        // Redirect to payment
        window.location.replace(data.url);
      })
      .catch((err) => {
        Notify.failure(`Error occured while trying to subscribe: ${err.message}`);
      });
  };
  const handlePassword = async (event: any) => {
    const data: oldPass = {
      email: user?.email,
      OldPassword: Password,
    };
    const dataUser: userResData = {
      email: user?.email,
      NewPassword: newPassword,
    };
    event.preventDefault();
    const passwordStrength = checkPasswordStrength(newPassword);
    const { request } = profileService.PostOldPassword(data);
    request
      .then((ress) => {
        setIsValiOld('is-valid');
        setIsError(false);
      })
      .catch((err) => {
        console.error(err);
        setIsValiOld('is-invalid');
        setIsError(true);
      });
    if (passwordStrength === 0) {
      setIsValiNew('is-invalid');
      setIsError(true);
    } else {
      setIsValiNew('is-valid');
      setIsError(false);
    }
    if (newPassword !== PasswordRepeat || PasswordRepeat === '') {
      setIsValiRep('is-invalid');
      setIsError(true);
    } else {
      setIsValiRep('is-valid');
      setIsError(false);
    }
    if (!isError) {
      console.log('log');
      const { request } = profileService.PostNewPassword(dataUser);
      request
        .then((res) => {
          Notify.success('You change your password');
          setNewPassword('');
          setPassword('');
          setPasswordRepeat('');
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };

  return (
    <div className="container mt-5 p-5 shadow-lg">
      <div className="row">
        <div className="mt-5  col-md-3">
          <FiUser size={270}></FiUser>
        </div>
        <div className="col-md-8">
          <h2>Welcome to my profile</h2>
          <div className="col-8">
            <div className="input-group input-group-sm mb-3">
              <span className="input-group-text" id="inputGroup-sizing-sm">
                Your e-mail
              </span>
              <input
                type="text"
                className="form-control"
                placeholder={user?.email}
                aria-label="Sizing example input"
                aria-describedby="inputGroup-sizing-sm"
              />
            </div>
            <div className="form-group border p-2">
              {user?.role !== 'free' ? (
                <>
                  <p className="text-success">Premium is active</p>
                  <div className="d-flex flex-row">
                    <strong className="mx-2">Valid From: </strong>
                    <span>{new Date(pack?.valid_from as string).toLocaleDateString()}</span>
                  </div>
                  <div>
                    <strong className="mx-2">Days left: </strong>
                    <span>{daysLeft}</span>
                  </div>
                </>
              ) : (
                <>
                  <p>
                    You don&apos;t have the premium package.
                    <br /> Would you like to buy it?
                  </p>
                  <button type="button" onClick={handleSubscribe} className="btn btn-warning">
                    Buy subscription
                  </button>
                </>
              )}
            </div>
          </div>
          <div className="text-center mt-3 float-start has-validation">
            <button type="button" onClick={() => navigate('/history')} className="btn btn-primary float-start">
              Look at your watch history
            </button>
            <br />
            <button
              type="button"
              onClick={() => {
                navigate('/Report');
              }}
              className="btn btn-danger float-start mt-3"
            >
              Report an Error
            </button>
            <br />

            <br />
          </div>
        </div>
        <div className=" col-md-3 text-center "></div>
        <form className=" col " onSubmit={handlePassword}>
          <div className=" col-md-3 text-center mt-3">
            <div className=" col text-center mt-3">
              Old password
              <input
                type="password"
                className={'form-control ' + isValiOld}
                placeholder="Old Password"
                aria-label="OldPassword"
                aria-describedby="basic-addon1"
                value={Password}
                onChange={(event) => setPassword(event.target.value)}
              />
              <div className="invalid-feedback">Please enter your old password</div>
            </div>
            <div className=" col text-center mt-3">
              New password
              <input
                type="password"
                className={'form-control ' + isValiNew}
                placeholder="New Password"
                aria-label="NewPassword"
                aria-describedby="basic-addon1"
                value={newPassword}
                onChange={(event) => setNewPassword(event.target.value)}
              />{' '}
              <div className="invalid-feedback">Your password is too weak </div>
            </div>
            <div className=" col text-center mt-3">
              Repeat password
              <input
                type="password"
                className={'form-control ' + String(isValiRep)}
                placeholder="Repeat password"
                aria-label="PasswordRepeat"
                aria-describedby="basic-addon1"
                value={PasswordRepeat}
                onChange={(event) => setPasswordRepeat(event.target.value)}
              />
              <div className="invalid-feedback"> The passwords are not identical</div>
              <br />
            </div>
            <button className={isError ? 'btn btn-danger' : 'btn btn-success'}>Reset password</button>
          </div>
        </form>
      </div>
    </div>
  );
};
