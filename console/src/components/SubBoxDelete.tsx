import React, { useState } from 'react';
import api from '../api';

type MyProps = {
	isProject: boolean;
	correctId: string;
};

const SubBoxDelete = (props: MyProps) => {
	const { isProject, correctId } = props;

	const [ id, setId ] = useState('');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ]: any = useState(false);
	const [ boxEditable, setBoxEditable ] = useState(false); // Check if Title, Emoji, or Desc. has changed

	return (
		<div className="SubBoxName-c-00 SearchBox-c-00 Theme-Box-Shadow">
			<h4 className="Flex-Row-Space-Between">
				Delete Object Permanently {isAsync && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<p className="SearchBox-Description-P-01">Permanently delete this Project/Blog from the Database</p>

			{boxEditable && (
				<input
					placeholder="ID of the Object"
					value={id}
					onChange={(e) => {
						setId(e.target.value.trim());
					}}
				/>
			)}

			{errorMsg && (
				<p
					style={{
						color: 'var(--danger-red-color)',
						padding: '0px 20px 10px 20px'
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
				{boxEditable ? (
					<React.Fragment>
						<button
							className="Theme-Btn-Red"
							onClick={async () => {
								setErrorMsg(false);

								if (id !== correctId) {
									setErrorMsg("ID doesn't match");
									return;
								}

								setIsAsync(true);
								try {
									await api.delete(
										`/private/${isProject ? 'project' : 'blog'}/delete/permanent/${id}`
									);
									setErrorMsg(false);
									setBoxEditable(false);
									window.location.href = '/admin';
								} catch (e) {
									setErrorMsg(e);
								}
								setIsAsync(false);
							}}
						>
							Delete Permanently
						</button>
						<button
							style={{ marginLeft: '12px' }}
							className="Theme-Btn-Grey"
							onClick={() => {
								setId('');
								setErrorMsg(false);
								setBoxEditable(false);
							}}
						>
							Cancel
						</button>
					</React.Fragment>
				) : (
					<button className="Theme-Btn-Red" onClick={() => setBoxEditable(true)}>
						Delete
					</button>
				)}
			</div>
		</div>
	);
};

export default SubBoxDelete;
