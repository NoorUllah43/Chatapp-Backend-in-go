
const SendMessage = () => {
    const data = new Date();
    const time = data.getHours() + ' : ' + data.getMinutes();
  return (
   <div className='w-auto h-auto flex flex-col items-end'>
      <div className='bg-[#4A7BF7] text-white w-auto h-auto inline-block p-2 rounded-t-xl rounded-s-xl '>
        <p className='w-auto h-auto inline-block'>Lorem ipsum dolor sit amet consectetur adipisicing elit.</p>
      </div>
      <div className='w-auto h-auto inline-block p-2'>
        <p className='w-auto h-auto inline-block'>{time}</p>
      </div>
    </div>

  )
}

export default SendMessage
