namespace Application.Contracts.Extention
{
    public class ExtentionResponse
    {
        public ExtentionResponse(IList<ExtentionDto> extention)
        {
            Extention = extention;
        }
        public IList<ExtentionDto> Extention { get; set; }
    }
}
