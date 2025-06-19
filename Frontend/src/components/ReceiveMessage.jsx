
const ReceiveMessage = () => {
    const data = new Date();
    const time = data.getHours() + ' : ' + data.getMinutes();
  return (
     <div className='w-auto h-auto flex flex-col items-start'>
      <div className='bg-gray-200 w-auto h-auto inline-block p-2 rounded-t-xl rounded-e-xl '>
        <p className='w-auto h-auto inline-block'>Lorem ipsum dolor sit amet consectetur adipisicing elit.</p>
      </div>
      <div className='w-auto h-auto inline-block p-2'>
        <p className='w-auto h-auto inline-block'>{time}</p>
      </div>
    </div>
  )
}

export default ReceiveMessage
