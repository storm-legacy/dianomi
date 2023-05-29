import React, { useContext, useEffect, useState } from 'react';
import { FiUser } from 'react-icons/fi';
import { AuthContext } from '../../context/AuthContext';
import profileService from '../../services/profile.service';
import { Notify } from 'notiflix';
import { Package } from '../../services/admin.service';
import { CanceledError } from 'axios';
import { redirect, useNavigate } from 'react-router-dom';

export const ProfilePage = () => {
  const { user } = useContext(AuthContext);
  const navigate = useNavigate();

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
          <div className="text-center mt-3 float-start">
            <button type="button" className="btn btn-primary float-start">
              Look at your watch history
            </button>
            <br />
            <button type="button" className="btn btn-primary float-start mt-3">
              lorem ipsum
            </button>
            <br />

            <br />
          </div>
        </div>
      </div>
    </div>
  );
};
