using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace FreeERP.service
{

    public struct OrderRecord
    {
        public String order_id {get; set;}
        public String createtime { get; set; }
        public String total { get; set; }
        public String total_fact { get; set; }
        public String remark { get; set; }
        public String status { get; set; }
        public String create_user_id { get; set; }
    }
    class SoldManager
    {
    }
}
