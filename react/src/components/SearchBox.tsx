import React, { useState } from 'react';
import api from '../api';

const SearchBox = () => {
	const [ msg, setMsg ] = useState('');
	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ]: any = useState(false);

	return (
		<div className="SubBoxName-c-00 SearchBox-c-00 Theme-Box-Shadow">
			<h4 className="Flex-Row-Space-Between">
				Search Box {isAsync && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<p className="SearchBox-Description-P-01">Clear All Cached Files, Update all Documents</p>

			<input
				style={{ marginBottom: '10px' }}
				placeholder="Delete all the cache? (Yes/No)"
				value={msg}
				onChange={(e) => setMsg(e.target.value)}
			/>

			{errorMsg && (
				<p
					style={{
						color: 'var(--danger-red-color)',
						padding: '0px 0px 10px 20px'
					}}
					className="SearchBox-Description-P-01"
				>
					Error: {errorMsg}
				</p>
			)}

			<div
				style={{
					width: 'calc(100% - 40px)',
					padding: '2px 20px 20px 20px'
				}}
				className="Flex-Row-Start"
			>
				<button
					className="Theme-Btn-Green"
					onClick={async () => {
						setErrorMsg(false);

						if (msg.toLowerCase() !== 'yes') {
							setMsg('');
							setErrorMsg('please write "Yes" in the input');
							return;
						}

						setIsAsync(true);

						try {
							await api.delete(`/private/clear_cache/all`);
							setMsg('');
						} catch (e) {
							setMsg('');
							setErrorMsg(e);
						}

						setIsAsync(false);
					}}
				>
					Clear All Cache
				</button>
			</div>
		</div>
	);
};

export default SearchBox;
