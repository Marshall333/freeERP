using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using FreeERP.service;
using System.Threading;
using System.Net.Sockets;
using System.Net;

using Newtonsoft.Json;


namespace FreeERP.View
{
    public partial class SoldManagerForm : Form
    {
        public const String Host = "http://106.12.132.28:8389";
        public const String QueryOrderRecordsURL = "/get_order_list";
        public delegate void UpdateInvoke(List<OrderRecord> orderRecords);
        public struct OrderRecordsReq
        {
            public SoldManagerForm form { get; set; }
            public String beginTime { get; set; }
            public String endTime { get; set; }
        }

        public SoldManagerForm()
        {
            InitializeComponent();
        }

        public void queryOrderRecords(object temp)
        {
            OrderRecordsReq req = (OrderRecordsReq)temp;

            string resp = BgWork.HttpPost(Host + QueryOrderRecordsURL, String.Format("begin_time={0}&end_time={1}&page=1&page_row=10",
                this.dateTimePicker1.Text, this.dateTimePicker2.Text));
            JavaScriptObject jsonObj = JavaScriptConvert.DeserializeObject<JavaScriptObject>(resp);


            if(jsonObj.ContainsKey("code") && (Int64)jsonObj["code"] == 0 && jsonObj.ContainsKey("datas"))
            {
                JavaScriptArray jList = (JavaScriptArray)jsonObj["datas"];
                if(jList.Count == 0)
                {
                    MessageBox.Show("没有记录!");
                    return;
                }
                List<OrderRecord> orderRecords = new List<OrderRecord>();
                for(int i = 0; i < jList.Count; i++)
                {
                    JavaScriptObject item = (JavaScriptObject)jList[i];
                    OrderRecord tt = new OrderRecord {order_id=(string)item["order_id"], createtime=(string)item["createtime"],
                        total=(string)item["total"],total_fact=(string)item["total_fact"],status=(string)item["status"],remark=(string)item["remark"],
                        create_user_id=(string)item["create_user_id"]};
                    orderRecords.Add(tt);
                }
                req.form.BeginInvoke(new UpdateInvoke(updateSoldRecords), orderRecords);
            } else
            {
                MessageBox.Show("错误!", "查询失败!");
            }
        }

        private void querySoldRecords(object sender, EventArgs e)
        {
            Thread thread = new Thread(queryOrderRecords);
            OrderRecordsReq req = new OrderRecordsReq();
            req.form = this;
            req.beginTime = this.dateTimePicker1.Text;
            req.endTime = this.dateTimePicker2.Text;
            thread.Start(req);
        }

        // 更新订单记录列表
        public void updateSoldRecords(List<OrderRecord> orderRecords)
        { 
            foreach(OrderRecord data in orderRecords)
            {
                DataGridViewRow row = new DataGridViewRow();
                DataGridViewTextBoxCell col1 = new DataGridViewTextBoxCell();
                col1.Value = data.order_id;
                row.Cells.Add(col1);
                DataGridViewTextBoxCell col2 = new DataGridViewTextBoxCell();
                col2.Value = data.createtime;
                row.Cells.Add(col2);
                DataGridViewTextBoxCell col3 = new DataGridViewTextBoxCell();
                col3.Value = data.total;
                row.Cells.Add(col3);
                DataGridViewTextBoxCell col4 = new DataGridViewTextBoxCell();
                col4.Value = data.total_fact;
                row.Cells.Add(col4);
                DataGridViewTextBoxCell col5 = new DataGridViewTextBoxCell();
                col5.Value = data.remark;
                row.Cells.Add(col5);
                DataGridViewTextBoxCell col6 = new DataGridViewTextBoxCell();
                col6.Value = data.create_user_id;
                row.Cells.Add(col6);
                this.dataGridView1.Rows.Add(row);
            }
        }
    }
}
