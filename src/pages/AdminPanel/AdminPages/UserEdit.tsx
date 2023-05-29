import React, { ChangeEvent, useEffect, useState } from 'react';
import Modal from 'react-modal';
import adminService, { UserEditData, Package, Tier } from '../../../services/admin.service';
import { Block, Notify } from 'notiflix';
import { FaTrash, FaPlus, FaArrowRight } from 'react-icons/fa';
import { CanceledError } from 'axios';

interface UserEditProps {
  isOpen: boolean;
  onRequestClose: () => void;
  userId: number;
}

Modal.setAppElement('#root');

const UserEdit: React.FC<UserEditProps> = ({ isOpen, onRequestClose, userId }) => {
  const timeRegex = /^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[+-]\d{2}:\d{2})?)Z$/;
  const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      height: '60dvh',
      width: '50dvh',
      transform: 'translate(-50%, -50%)',
    },
  };

  const [email, setEmail] = useState('');
  const [verify, setVerify] = useState(false);
  const [reset, setReset] = useState(false);
  const [packages, setPackages] = useState<Package[]>([]);

  useEffect(() => {
    if (userId === 0) return;
    const { request, cancel } = adminService.getUser(Number(userId));
    request
      .then(({ data }) => {
        setEmail(data?.email);
        setVerify(data?.verified);
        setReset(false);
      })
      .catch((err) => {
        if (err instanceof CanceledError) return;
        Notify.failure('Could not fetch information about user');
        onRequestClose();
      });
    return () => cancel();
  }, [userId]);

  useEffect(() => {
    if (userId === 0) return;

    const { request, cancel } = adminService.getUserPackages(Number(userId));
    request
      .then(({ data }) => {
        setPackages(data);
      })
      .catch((err) => {
        if (err instanceof CanceledError) return;
        Notify.failure('Could not fetch user packages');
        onRequestClose();
      });

    return () => cancel();
  }, [userId]);

  const handleUserPatch = (email: string, reset: boolean, verified: boolean) => {
    const { request } = adminService.patchUser(userId, {
      email: email,
      reset_password: reset,
      verified: verified,
    });

    request
      .then((res) => {
        Notify.success('Changes were applied!');
      })
      .catch((err) => {
        Notify.failure(`Changes could not be applied: ${err.message}`);
      });

    onRequestClose();
    window.location.reload();
  };

  const handlePackageChange = (
    pack: Package,
    valid_from: string,
    valid_until: string,
    newTier: string,
    deletePackage = false,
  ) => {
    const validFrom = timeRegex.test(valid_from) ? valid_from : `${valid_from}T00:00:00Z`;
    const validUntil = timeRegex.test(valid_until) ? valid_until : `${valid_until}T00:00:00Z`;

    const tier = newTier == 'administrator' ? Tier.administrator : newTier == 'premium' ? Tier.premium : Tier.free;

    if (!deletePackage) {
      const { request } = adminService.patchUserPackage({
        id: pack.id,
        tier: tier,
        created_at: new Date(),
        user_id: userId,
        valid_from: validFrom,
        valid_until: validUntil,
      });
      request
        .then(() => {
          Notify.success('Values updated');
        })
        .catch((err) => {
          Notify.failure(`Could not update value: ${err.message}`);
        });
    } else {
      const { request } = adminService.deleteUserPackage(pack.id);
      request
        .then(() => {
          setPackages([]);
          Notify.success('Package successfuly removed!');
        })
        .catch((err) => {
          Notify.failure(`Package could not be removed: ${err.message}`);
        });
    }
  };

  const handleAddPackage = () => {
    const date = new Date();

    if (packages.length == 0) {
      const pack = {
        id: -1,
        user_id: userId,
        valid_from: date.toISOString(),
        valid_until: new Date(date.setDate(date.getDate() + 30)).toISOString(),
        tier: Tier.free,
        created_at: date,
      };
      const { request } = adminService.postUserPackage(pack);
      request
        .then(() => {
          setPackages([pack]);
          Notify.success(`Package successfully added!`);
        })
        .catch((err) => {
          Notify.failure(`Package could not be added: ${err.message}`);
        });
    }
  };

  return (
    <Modal isOpen={isOpen} onRequestClose={onRequestClose} contentLabel="Modal" style={customStyles}>
      <div className="content p-0" id="popup-window">
        <form
          style={{ margin: '25px' }}
          onSubmit={(e) => {
            e.preventDefault();
            handleUserPatch(email, reset, verify);
          }}
        >
          <h2>User edit</h2>
          <span>Email</span>
          <div className="p-1">
            <label>
              <input
                className="form-control"
                type="text"
                placeholder={email}
                value={email}
                onChange={(e) => {
                  setEmail(e.target.value);
                }}
              />
            </label>
          </div>
          <div className="p-1">
            <label className="form-check-label">
              <input
                className="form-check-input"
                type="checkbox"
                id="Very"
                checked={verify}
                onChange={(e) => {
                  setVerify(e.target.checked);
                }}
              />{' '}
              Verified
            </label>
          </div>
          <div className="p-1">
            <label className="form-check-label">
              <input
                className="form-check-input"
                type="checkbox"
                checked={reset}
                onChange={(e) => {
                  setReset(e.target.checked);
                }}
              />{' '}
              Send reset link
            </label>
          </div>
          <div className="d-flex flex-row justify-content-center py-2">
            <button type="submit" className="btn btn-success mx-2">
              Save
            </button>
            <button
              type="button"
              className="btn btn-secondary mx-2"
              onClick={() => {
                onRequestClose();
              }}
            >
              Cancel
            </button>
          </div>
        </form>
        <div className="row my-2">
          {packages.map((p, index) => {
            return (
              <div key={index} className="col-12 border p-2">
                <div className="d-flex justify-content-between mb-2">
                  <div>
                    <b>Pakiet:</b>
                    <select
                      className="mx-4"
                      onChange={(e) =>
                        handlePackageChange(p, String(p.valid_from), String(p.valid_until), e.target.value)
                      }
                    >
                      <option value="free">Free</option>
                      <option value="administrator">Administrator</option>
                      <option value="premium">Premium</option>
                    </select>
                  </div>
                  <button
                    type="button"
                    onClick={() => handlePackageChange(p, '', '', p.tier, true)}
                    className="btn btn-danger py-0 px-1"
                  >
                    <FaTrash></FaTrash>
                  </button>
                </div>
                <div className="d-flex flex-row">
                  <input
                    type="date"
                    id="validFrom"
                    value={String(p.valid_from).slice(0, 10)}
                    onChange={(e) => {
                      handlePackageChange(p, e.target.value, String(p.valid_until), p.tier);
                    }}
                  />
                  <span>
                    <FaArrowRight></FaArrowRight>
                  </span>
                  <input
                    type="date"
                    id="validTo"
                    value={String(p.valid_until).slice(0, 10)}
                    onChange={(e) => {
                      handlePackageChange(p, String(p.valid_from), e.target.value, p.tier);
                    }}
                  />
                </div>
              </div>
            );
          })}
          <button className="btn btn-primary" onClick={handleAddPackage} type="button">
            <FaPlus></FaPlus>
          </button>
        </div>
      </div>
    </Modal>
  );
};

export default UserEdit;
