import React from 'react';

interface props {
  postsPerPage: number;
  totalPosts: number;
  paginate: (num: number) => void;
}

const Paginate = ({ postsPerPage, totalPosts, paginate }: props) => {
  const pageNumbers: number[] = [];

  for (let i = 1; i <= Math.ceil(totalPosts / postsPerPage); ++i) pageNumbers.push(i);

  return (
    <div className="pagination-container">
      <nav>
        <ul className="pagination">
          <li className="page-item px-1">
            <button
              className="page-link"
              onClick={() => {
                paginate(pageNumbers[0]);
              }}
              aria-label="Previous"
            >
              <span aria-hidden="true">&laquo;</span>
            </button>
          </li>
          {pageNumbers.map((num) => (
            <li key={num} className="page-item px-1">
              <button
                className="page-link"
                onClick={() => {
                  paginate(num);
                }}
              >
                {num}
              </button>
            </li>
          ))}
          <li className="page-item px-1">
            <button
              className="page-link"
              onClick={() => {
                paginate(pageNumbers[pageNumbers.length - 1]);
              }}
              aria-label="Next"
            >
              <span aria-hidden="true">&raquo;</span>
            </button>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Paginate;
