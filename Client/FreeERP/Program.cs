using FreeERP.View;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;

namespace FreeERP
{
    static class Program
    {
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>

        static MainPage mainPage;

        [STAThread]
        static void Main()
        {
            Program.mainPage = new MainPage();


            Application.EnableVisualStyles();
            //Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(Program.mainPage);
        }
    }
}
